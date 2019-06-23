pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                // sh 'echo skip build'
                sh 'go get github.com/ogier/pflag'
                sh 'go build main.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'go get github.com/ogier/pflag'
                sh 'go test -v'
                sh 'go test -bench=.'
            }
        }
        stage('Lint') {
            agent {
                // docker { image 'obraun/vss-jenkins' }
                docker { image 'obraun/vss-protoactor-jenkins' }
            }   
            steps {
                sh 'go get github.com/ogier/pflag'
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
