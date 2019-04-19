FastTrack-Golang-Quiz

This is my first golang project which serves has a solution to https://gist.github.com/Tazer/fd02311a143afac097dc4dfaa47667f6

Installation Instructions

1) "go install" in project to install quiz-cli command 
2) Start the restful api server by running "quiz-cli server"
3) To Start the CLI client run "quiz-cli client"

Docker Installation Instructions

1) Build Container by running "sudo docker build -t quiz-app ."
2) Run server inside docker container using "sudo docker run -it -p 3001:3001 quiz-app server"
3) On a separate terminal run the quiz client  by "sudo docker run -it --network="host" quiz-app client" 

Info
The project is using cobra has a CLI Framework has suggested and its own http API.

You can check out some of my personal and school related projects at my personal git repo hosted on: http://git.scifo.org/explore/repos

I have been working in the IT industry for the last 4 years were I primarily worked with PHP(Laravel) , and Angular .
I have also spent some time developing android Apps on a professional level.I Really enjoyed golang during the development of this small project and hope to do it again.
I'm also a big fan of linux and opensource, I do have hands on real experience in linux as during my last job it was my responsibility to mange and overlook various linux VPSes. 