#!groovy
@Library('github.com/cloudogu/ces-build-lib@2.1.0')
import com.cloudogu.ces.cesbuildlib.*

// Creating necessary git objects, object cannot be named 'git' as this conflicts with the method named 'git' from the library
gitWrapper = new Git(this, "cesmarvin")
gitWrapper.committerName = 'cesmarvin'
gitWrapper.committerEmail = 'cesmarvin@cloudogu.com'
gitflow = new GitFlow(this, gitWrapper)
github = new GitHub(this, gitWrapper)
changelog = new Changelog(this)
gpg = new Gpg(this, docker)
Makefile makefile = new Makefile(this)

// Configuration of repository
repositoryOwner = "cloudogu"
repositoryName = "ces-control-api"

// Configuration of branches
productionReleaseBranch = "main"
developmentBranch = "develop"
currentBranch = "${env.BRANCH_NAME}"

node('docker') {
    timestamps {
        properties([
                // Keep only the last x builds to preserve space
                buildDiscarder(logRotator(numToKeepStr: '10')),
                // Don't run concurrent builds for a branch, because they use the same workspace directory
                disableConcurrentBuilds(),
        ])

        stage('Checkout') {
            checkout scm
        }

        if (gitflow.isReleaseBranch()) {
            String releaseVersion = gitWrapper.getSimpleBranchName()
            String version = makefile.getVersion()

            stage('Finish Release') {
                gitflow.finishRelease(releaseVersion, productionReleaseBranch)
            }

            stage('Add Github-Release') {
                releaseId = github.createReleaseWithChangelog(releaseVersion, changelog, productionReleaseBranch)
            }
        }
    }
}
