#! /bin/sh
#
# post.sh
# Copyright (C) 2018 MacBookPro <MacBookPro@MacBooks-MBP-5>
#
# Distributed under terms of the MIT license.
#


curl -X POST \
  -H "Authorization: Key MyKey " \
  -H "Content-Type: application/json" \
  -d '
    {
      "inputs": [
      {
          "data": {
              "image": {
                "url":
                "https://c1.staticflickr.com/4/3837/14230768839_436bc3cf5e_o.jpg"
                }
            }
        }
    ]
  }'\
    https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs
