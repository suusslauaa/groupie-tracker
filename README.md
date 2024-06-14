## groupie-tracker

### Objectives

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- [API](https://groupietrackers.herokuapp.com/api) consists in four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

### Instructions

- open terminal with `Ctrl + Alt + T`
- clone the repo with this command `git clone git@git.01.alem.school:aserikbay/groupie-tracker.git`
```console
student@ALEM:~$ git clone git@git.01.alem.school:aserikbay/groupie-tracker.git
```
- enter the folder with command `cd groupie-tracker/`
```console
student@ALEM:~$ cd groupie-tracker/
```
- run the program `go run .`
```console
student@ALEM:~$ go run .
```
- if you want change the port add `-addr :PORT` flag, like this `go run . -addr :8080` 
- follow the [link](http://localhost:4000/) on terminal
