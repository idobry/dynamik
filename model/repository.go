package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

//var BaseDir = viper.GetString("base-dir")
const BaseDir = "./workspace/"
var repo Repository

type Repository struct{
	git *git.Repository
	repoName string
	url string
	username string
	password string
	token string
}

func NewRepository() *Repository{
	var url = viper.GetString("git-url")
	repoName := strings.Split(url[strings.LastIndex(url, "/")+1:], ".")[0]
	repo = Repository{repoName: repoName, url: viper.GetString("git-url"), password: viper.GetString("password"), token: viper.GetString("token")}

	if repo.token != ""{
		repo.git = token()
	}else if repo.password != ""{
		repo.git = usernamePassword()
	}else if repo.username == ""{
		repo.git = publicRepo()
	}else{
		logrus.Error("Wrong git config input")
	}

	return &repo
}

func publicRepo() *git.Repository{
	r, err := git.PlainClone(BaseDir+"/"+repo.repoName, false, &git.CloneOptions{
		URL:               repo.url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil{
		panic(err)
	}

	return r
}

func token() *git.Repository{
	r, err := git.PlainClone(BaseDir+"/"+repo.repoName, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "username",
			Password: repo.token,
		},
		URL:      repo.url,
		Progress: os.Stdout,
	})

	if err != nil{
		panic(err)
	}

	return r
}

func usernamePassword() *git.Repository{
	r, err := git.PlainClone(BaseDir+"/"+repo.repoName, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: repo.username,
			Password: repo.password,
		},
		URL:      repo.url,
		Progress: os.Stdout,
	})

	if err != nil{
		panic(err)
	}

	return r
}
