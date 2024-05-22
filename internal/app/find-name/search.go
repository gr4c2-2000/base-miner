package findname

import (
	"strings"
	"sync"
)

type Search struct {
	nameMap          map[string]struct{}
	surnameMap       map[string]struct{}
	jobChan          chan *Job
	concurrentTasks  int
	concurrentSearch int
	resultChan       chan *Person
}

type Job struct {
	Url     string
	Content string
	end     chan struct{}
}

type Person struct {
	Source  string
	Name    []string
	Surname []string
}

func (s *Search) Close() {
	close(s.jobChan)
	close(s.resultChan)
}

func (s *Search) PushJob(url, content string) chan struct{} {
	end := make(chan struct{}, 1)
	s.jobChan <- &Job{Content: content, Url: url, end: end}
	return end
}

func (s *Search) Result() chan *Person {
	return s.resultChan
}

func NewSearch(concurrentTasks, concurrentSearch int) *Search {
	search := &Search{
		nameMap:          ParseNamesToMap(),
		surnameMap:       ParseSurNameToMap(),
		jobChan:          make(chan *Job, concurrentTasks),
		resultChan:       make(chan *Person, concurrentTasks*concurrentSearch),
		concurrentTasks:  concurrentTasks,
		concurrentSearch: concurrentSearch,
	}

	for range concurrentTasks {
		go func() {
			search.workerTasks()
		}()
	}

	return search
}

func (s *Search) workerTasks() {
	for j := range s.jobChan {
		if s.concurrentSearch == 0 {
			s.inText(j)
			continue
		}
		s.inTextAsync(j)
	}
}

func (s *Search) inText(job *Job) {
	text := s.prepareWords(job.Content)
	for index, _ := range text {
		person := s.MatchFullName(index, text)
		if person != nil {
			s.resultChan <- person
		}
	}
	close(job.end)
}

func (s *Search) inTextAsync(job *Job) {
	var (
		matchChan = make(chan int, s.concurrentSearch)
		wg        = sync.WaitGroup{}
		text      = s.prepareWords(job.Content)
	)

	for range s.concurrentSearch {
		wg.Add(1)
		go func() {
			for index := range matchChan {
				person := s.MatchFullName(index, text)
				if person != nil {
					s.resultChan <- person
				}
			}
			wg.Done()
		}()
	}

	for idx, _ := range text {
		matchChan <- idx
	}

	close(matchChan)
	wg.Wait()
	close(job.end)
}

func (s *Search) MatchFullName(index int, strSlice []string) *Person {
	p := &Person{}
	if !s.MatchName(strSlice[index], p) {
		return nil
	}
	s.MatchRecursive(index, strSlice, p, 1)
	s.MatchRecursive(index, strSlice, p, -1)
	return p
}

func (s *Search) MatchRecursive(index int, strSlice []string, person *Person, direction int) {
	index = index + direction
	if index < 0 || index == len(strSlice) {
		return
	}
	if !s.MatchStrategy(strSlice[index], person) {
		return
	}
	s.MatchRecursive(index, strSlice, person, direction)

}

func (s *Search) MatchStrategy(world string, person *Person) bool {
	switch {
	case s.MatchName(world, person):
	case s.MatchSurname(world, person):
	default:
		return false
	}
	return true
}

func (s *Search) MatchSurname(name string, person *Person) bool {
	if len(name) < 2 {
		return false
	}
	name = strings.ToUpper(name)
	if _, ok := s.surnameMap[name]; ok {
		person.Surname = append(person.Surname, name)
		return true
	}
	return false
}

func (s *Search) MatchName(name string, person *Person) bool {
	if len(name) < 2 {
		return false
	}
	name = strings.ToUpper(name)
	if _, ok := s.nameMap[name]; ok {
		person.Name = append(person.Name, name)
		return true
	}
	return false
}

func (s *Search) prepareWords(str string) []string {
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, ":", "")
	str = strings.ReplaceAll(str, ";", "")
	str = strings.ReplaceAll(str, "!", "")
	str = strings.ReplaceAll(str, "?", "")
	str = strings.ReplaceAll(str, "\n", " ")
	str = strings.ReplaceAll(str, "\r", " ")
	str = strings.ReplaceAll(str, "\r", " ")
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "[", " ")
	str = strings.ReplaceAll(str, "]", " ")
	str = strings.ReplaceAll(str, "(", " ")
	str = strings.ReplaceAll(str, ")", " ")
	str = strings.ReplaceAll(str, "{", " ")
	str = strings.ReplaceAll(str, "}", " ")
	words := strings.Split(str, " ")
	return words
}
