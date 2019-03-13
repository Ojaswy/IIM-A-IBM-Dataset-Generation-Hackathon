Hey,
   I am Ojaswy Akella, a major data science enthusiast. I am currently pursuing Integrated MTech at University of Hyderabad. This text file holds the explanation for my solution to the IIMA-IBM Dataset Generation Hackathon. I approached the problem using the language Go, because it was taught to me at our University and I had some exprerince by solving a few assignments in 'Theory of Computation' and I felt this approach was much better than the one with Python.

The results are saved in the file 'output.json' 

My solution includes the use of Regular Languange and the corresponding 'regexp' library in Go. Using RE to identify the pattern in strings and replacing them with their corresponding entities was the basic idea. I have used the "github.com/Pallinder/go-randomdata" and "github.com/brianvoe/gofakeit" repositories as datasets for entitites.

Implementation:

1) Prerequisites: Git
2) go get gopkg.in/jdkato/prose.v2
3) go get "github.com/Pallinder/go-randomdata"
4) go get "github.com/brianvoe/gofakeit"
5) go build

After the final step, An application pops up. Run that Application till the process completes, after which the output file will be generated.

References: [](https://golang.org/pkg/regexp/)

