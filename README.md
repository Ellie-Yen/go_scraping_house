# go_scraping_house

## Purpose

For my personal usage only.
Looking for rental houses that located in taipei in a specific budget range & exclude keywords.

## file structure

The file structure will be changed over time.

```

├── api/                        # the scraping source api
  ├── common.go                 # a wrapper for http request since I'm lazy to write headers manaully
  ├── constants.go              # the fix part of request header / cookies
  ├── houseList.go              # the endpoint to fetch a house list with certain condition
  ├── responses.go              # stores the response structure
  ├── main.go                   # the entry point that handle batch fetch house lists

├── controllers/                # route controllers
├── routes/                     # gin router set-ups
├── templates/                  # gin frontend html templates
├── utils/

├── main.go                     # entrypoint

├── go.mod
├── go.sum
├── README.md
├── .gitignore
```
