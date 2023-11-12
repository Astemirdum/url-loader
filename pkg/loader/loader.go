package loader

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Astemirdum/url-loader/pkg"
	"golang.org/x/sync/errgroup"
)

type loadURLResult struct {
	u    string
	size int
	ltv  time.Duration
}

type Processor struct {
	glimit int
}

func NewProcessor(glimit int) *Processor {
	return &Processor{
		glimit: glimit,
	}
}

func (p *Processor) Run(reader io.Reader) error {
	urlCh := make(chan string)
	sc := bufio.NewScanner(reader)
	go func() {
		defer close(urlCh)
		for sc.Scan() {
			u := strings.TrimSpace(sc.Text())
			urlCh <- u
		}
	}()

	wgDone := sync.WaitGroup{}
	errCh := make(chan error)
	wgDone.Add(1)
	go func() {
		defer wgDone.Done()
		for err := range errCh {
			fmt.Fprintf(os.Stderr, "process url err: %v\n", err)
		}
	}()

	resCh := make(chan loadURLResult)
	wgDone.Add(1)
	go func() {
		defer wgDone.Done()
		for res := range resCh {
			fmt.Printf("process url result: (url: %s, size: %d bytes, ltv: %s)\n", res.u, res.size, res.ltv.String())
		}
	}()

	gg := errgroup.Group{}
	ctx := context.Background()
	for i := 0; i < p.glimit; i++ {
		gg.Go(func() error {
			cl := &http.Client{Timeout: time.Minute}
			for u := range urlCh {
				res, err := loadURL(ctx, u, cl)
				if err != nil {
					errCh <- err
				} else {
					resCh <- res
				}
			}

			return nil
		})
	}
	go func() {
		defer func() {
			close(errCh)
			close(resCh)
		}()
		if err := gg.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "something went wrong '%s'", err)
		}
	}()

	wgDone.Wait()
	return nil
}

func loadURL(ctx context.Context, u string, client *http.Client) (loadURLResult, error) {
	dur := time.Now()
	if err := pkg.ValidateURL(u); err != nil {
		return loadURLResult{}, fmt.Errorf("url=%s err=%w", u, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, http.NoBody)
	if err != nil {
		return loadURLResult{}, fmt.Errorf("new req url=%s err=%w", u, err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return loadURLResult{}, fmt.Errorf("client.Do(req) url=%s err=%w", u, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return loadURLResult{}, fmt.Errorf("io.ReadAll url=%s err=%w", u, err)
	}

	return loadURLResult{
		u:    u,
		size: len(data),
		ltv:  time.Since(dur),
	}, nil
}
