# locator

Performance enhanced, vector based suggestion api.

### How it works
1. Loops through locations dataset and creates a vector based dataset of all the locations, making them easy for the CPU to process in long-run.
2. These vectors are then used to find the closest location to the user's location.
3. The api will provide location based suggestions considering the input from user.

### Demo
![Gif](./art/golocator.gif)

![vector based dataset](./art/vector_dataset.png)


#### Dataset
https://data.opendatasoft.com/api/explore/v2.1/catalog/datasets/geonames-postal-code@public/exports/json?lang=en&timezone=Asia%2FKolkata


place -> city, state, country, zip code
abc
adcd 
abcde 

---
LICENSE

<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><a property="dct:title" rel="cc:attributionURL" href="https://github.com/fayaz07/locator">locator</a> by <a rel="cc:attributionURL dct:creator" property="cc:attributionName" href="https://github.com/fayaz07">Mohammad Fayaz</a> is licensed under <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">Attribution-NonCommercial-ShareAlike 4.0 International<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/sa.svg?ref=chooser-v1"></a></p>
