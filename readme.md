# go-module1

## Motivation

Right now we are migrating from Travis to Github as CI for some reasons out of the scope of this task.

This migration can require some changes in source code (repository). The goal is to minimize this effort now keeping in mind any following migration.

**Note:** There is a project that helps to run github actions locally https://github.com/nektos/act but I think that it is over-engeening after taking a look on it. 
**Note 2:** There is also https://docs.gitlab.com/runner/ that allows to run Github actions in our infraestructure but it also ends with more complexity in our devops and developer teams.

Have a look at this question on StackOverflow. https://stackoverflow.com/questions/59241249/how-to-run-github-actions-workflows-locally?answertab=votes#tab-top


## Perfect example:
As a perfect example about minimizing effort in the migration could be a project in Java using Gradle as build tool.

JDK environment is required. 

Gradle plugins are used to do actions in our project without installing new software.
They are similar to Github actions in the way that they are third-party tools that no installing is required.

Dependencies between tasks came out of the box.

Every event on Github can be mapped to a single gradle command call.
 ./gradlew clean build test sonar publish

Gradle commands run everywhere.

Gradle also accepts running shell scripts, as Github actions do.

If we avoid to do not use specific CI features migrations will be easier and reproducing what it is running in the CI will be also easy.

## Case studies for GoLang projects
https://www.baeldung.com/ant-maven-gradle
### Maven 
It is related to Java and difficult to customize because it is based on the principle of conventions over configuration that could not fit very well in our case.

### Ant
It is based on an XML file and everything should be declared on it. It has also plugins and it is very configurable. It needs also JRE environment. 

### Gradle
It is a build tool that admits other programming languages. It is versatile and there is a go plugin but it seems to be outdated. Tasks dependencies is also handled and there are lots of plugings.

### Make
It is one of the oldest tools. It also helps with task dependencies. Perfect for use with shell scripts. Tools should be installed manually because there is no concept of plugin.

## Finally
This repo is just an starting point for using gradle and GoLang together to see if it fits to us.