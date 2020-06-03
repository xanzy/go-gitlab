package main

import (
	"fmt"
	"log"
  "sync"
  "time"

	"github.com/xanzy/go-gitlab"
  "github.com/negah/percent"
)

var waitGroup sync.WaitGroup

func main() {
	git, err := gitlab.NewClient(
    "token",
    gitlab.WithBaseURL("url"),
	)
	if err != nil {
		log.Fatal(err)
	}

	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
			Page:    1,
		},
	}

  // var languageSlices []*gitlab.ProjectLanguages
	var metrics []metric
  var count = 0

	for {
		// Get the first page with projects.
		ps, resp, err := git.Projects.ListProjects(opt)
    
		if err != nil {
			log.Fatal(err)
		}
    waitGroup.Add(len(ps))

		// List all the projects we've found so far.
		for _, p := range ps {
      count++
      go func(project *gitlab.Project) {
        defer waitGroup.Done()
        languages, _, err := git.Projects.GetProjectLanguages(
          project.PathWithNamespace,
        )
        if err != nil {
          log.Fatal(err)
        }
        
        for key, _ := range *languages {

          has := hasElement(metrics, key)
          if !has {
            metrics = append(
              metrics,
              metric{Language: key, Count: 1},
            )
          } else {
            increaseElement(metrics, key, count)
          } 
          fmt.Printf(".")
          time.Sleep(5e7)
        }
      }(p)
    }

    waitGroup.Wait()
    
		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
      for _, metric := range metrics {
        fmt.Printf(
          "\n%s %d %0.2f%%",
          metric.Language,
          metric.Count,
          metric.Percentage,
        )
      }
      fmt.Printf("\n\ntotal projects found: %d\n", count)
      fmt.Println("_____________________________")
      break
    }

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
}

func hasElement(slice []metric, key interface{}) bool {
  for i := 0; i < len(slice); i++ {
    if key == slice[i].Language {
      return true
    }
  }
  return false
}

func increaseElement(
  slice []metric,
  key interface{},
  counter int,
) {
  for i := 0; i < len(slice); i++ {
    if key == slice[i].Language {
      slice[i].Count++
      slice[i].Percentage = percent.PercentOf(
        slice[i].Count,
        counter,
      )
    }
  }
}


type metric struct {
  Language string 
  Count int
  Percentage float64
}
