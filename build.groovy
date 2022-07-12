pipeline {
    agent any
    parameters{
        string(name: "BRANCH_NAME", defaultValue: "siva")
        string(name: "url", defaultValue: "https://github.com/sivalakshmanna/cicd-demo-golang.git")
    }
    
    stages{
        stage("chekout code"){
            steps{
                println "clone our code to our repository"
                sh "ls -l"
                
                

            }
        }
    
        stage("build code"){
            steps{
                println "go build"
                sh "go build main.go"
                
            }
        }
    }
}