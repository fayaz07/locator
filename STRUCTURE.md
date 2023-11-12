# locator structure

## Modules

### 1. CLI module for preparing dataset
The role of this module is to prepare the dataset for the api to use. The dataset is prepared in the form of a vector based dataset. The dataset is prepared by looping through the dataset and creating a vector based dataset of all the locations, making them easy for the CPU to process in long-run.

### 2. Core module
This module will contain all the common code that will be used by the api and the cli module, such as data models, json and csv parsers, etc.
