#!/usr/bin/env python
import requests
from data import CATEGORIES, PERFORMERS, DIRECTORS, STUDIOS, SCENES, SERIES
from slugify import slugify
import json

HOST = "http://127.0.0.1:3000"
COLLECTION_PATH = "%s/v1/collection" % HOST


def categories():

    for category in CATEGORIES:

        payload = {
            "name": category,
            "slug": slugify(category)
        }
        path = "%s/category" % COLLECTION_PATH
        print(path)
        r = requests.post(url=path, data=json.dumps(payload))

        print r.text


def performers():
    for performer in PERFORMERS:

        path = "%s/performer" % COLLECTION_PATH

        r = requests.post(url=path, data=json.dumps(performer))
        print r.text


def directors():
    for director in DIRECTORS:

        path = "%s/performer" % COLLECTION_PATH

        r = requests.post(url=path, data=json.dumps(director))
        print r.text


def studios():
    for studio in STUDIOS:

        path = "%s/studio" % COLLECTION_PATH

        r = requests.post(url=path, data=json.dumps(studio))
        print r.text


def scenes():
    for studio in SCENES:

        path = "%s/scene" % COLLECTION_PATH

        r = requests.post(url=path, data=json.dumps(studio))
        print r.text


def series():
    for series in SERIES:

        path = "%s/series" % COLLECTION_PATH

        r = requests.post(url=path, data=json.dumps(series))
        print r.text


def main():
    categories()
    performers()
    directors()
    studios()
    scenes()
    series()


if __name__ == '__main__':
    main()
