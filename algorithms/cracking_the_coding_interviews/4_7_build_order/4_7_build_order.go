package main

import "fmt"

type Project struct {
	name         string
	dependencies []Project
}

func (p *Project) getNumDependencies() int {
	return len(p.dependencies)
}

func removeDependency(dependOn Project, projects *[]*Project) {
	for _, p := range *projects {
		indexToRemove := -1
		for i, d := range p.dependencies {
			if d.name == dependOn.name {
				indexToRemove = i
			}
		}
		if indexToRemove != -1 {
			part1 := p.dependencies[:indexToRemove]
			part2 := p.dependencies[indexToRemove+1:]

			var remainDependencies []Project
			remainDependencies = append(remainDependencies, part1...)
			remainDependencies = append(remainDependencies, part2...)
			p.dependencies = remainDependencies
		}
	}
}

func findBuildOrder(projects *[]*Project) []Project {
	var projectsOrder []Project
	visited := make(map[string]bool)
	for _, p := range *projects {
		if p.getNumDependencies() == 0 {
			visited[p.name] = true
			projectsOrder = append(projectsOrder, *p)
		}
	}

	indexToProcess := 0
	for indexToProcess < len(projectsOrder) && indexToProcess < len(*projects) {
		current := projectsOrder[indexToProcess]

		removeDependency(current, projects)

		for _, p := range *projects {
			if p.getNumDependencies() == 0 && !visited[p.name] {
				visited[p.name] = true
				projectsOrder = append(projectsOrder, *p)
			}
		}

		indexToProcess++
	}

	if len(projectsOrder) == len(*projects) {
		return projectsOrder
	} else {
		return nil
	}
}

func main() {
	projects := &[]*Project{
		&Project{name: "a"},
		&Project{name: "b"},
		&Project{name: "c"},
		// &Project{name: "d"},
		// &Project{name: "e"},
		// &Project{name: "f"},
		// &Project{name: "g"},
		// &Project{name: "h"},
		// &Project{name: "k"},
	}
	p := *projects
	p[0].dependencies = []Project{*p[1]}
	p[2].dependencies = []Project{*p[0]}
	p[1].dependencies = []Project{*p[2]}
	// p[6].dependencies = []Project{*p[2], *p[0], *p[3]}
	// p[7].dependencies = []Project{*p[6]}
	// p[5].dependencies = []Project{*p[8]}

	res := findBuildOrder(&p)

	if res == nil {
		fmt.Println("no possible order!")
		return
	}

	for _, r := range res {
		fmt.Println(r.name)
	}
}
