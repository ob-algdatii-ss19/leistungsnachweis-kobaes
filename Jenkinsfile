pipeline {
    agent none
    stages {
        stage('Test') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }
            steps {
                sh 'echo go test -v'
                sh 'echo go test -bench=.'
                sh 'cd sorting && go get -v -d -t ./...'
                sh 'go get github.com/t-yuki/gocover-cobertura' // install Code Coverage Tool
                sh 'cd sorting && go test -v -coverprofile=cover.out’ // save coverage info to file
                sh 'gocover-cobertura < sorting/cover.out > coverage.xml’ // transform coverage info to jenkins readable format
                sh 'cd sorting && go test -bench=.'
                publishCoverage adapters: [coberturaAdapter('coverage.xml’)] // publish report on Jenkins
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }   
            steps {
                sh 'golangci-lint run --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent {
                label 'master'
            }
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME}"
            }
        }
    }
    post {
        changed {
            script {
                if (currentBuild.currentResult == 'FAILURE') { // Other values: SUCCESS, UNSTABLE
                    // Send an email only if the build status has changed from green/unstable to red
                    emailext subject: '$DEFAULT_SUBJECT',
                        body: '$DEFAULT_CONTENT',
                        recipientProviders: [
                            [$class: 'DevelopersRecipientProvider']
                        ], 
                        replyTo: '$DEFAULT_REPLYTO'
                }
            }
        }
    }
}
