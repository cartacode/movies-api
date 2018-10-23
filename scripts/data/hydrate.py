#!/usr/bin/env python
import time

import requests
import json
from slugify import slugify
import os
import pprint
pp = pprint.pprint

HOST = "https://api-stage.vuli.tv"
# HOST = "http://127.0.0.1:3000"


def openDataFile(path):

    data = None

    with open(path,) as fileh:
        data = json.load(fileh)
    return data


def category(data):
    for line in data:
        r = requests.post('%s/category' % HOST,
                          json={
                              'title': line['Category_Name'],
                              'slug': slugify(line['Category_Name'])})
        print r.text


def directors(data):
    for line in data:
        r = requests.post('%s/performer' % HOST,
                          json={
                              'name': line['Name'],
                              'slug': slugify(line['Name'])})
        print r.json()


def loadToCache(path, filename):

    if os.stat("data/%s.json" % filename):
        print "Cache exists"
        return

    jsonfiles = []
    for root, dirs, files in os.walk(path):
        # path = root.split(os.sep)
        for file in files:
            if 'json' in file:
                try:
                    data = openDataFile(
                        "%s/%s" % (root, file))
                    # pp(data)

                except Exception as e:
                    print(e)
                jsonfiles.append(data)

    with open("%s.json" % filename, "w") as movie:
        movie.write(json.dumps(jsonfiles))


def length(timedelta):
    timedelta = timedelta.split(':')
    length = 0
    length += int(timedelta[0])*60*60
    length += int(timedelta[1])*60
    length += int(timedelta[2])
    return length


def findObjectID(uri, slug):
    # print slugify(slug)
    r = requests.get('%s/%s/slug/%s' % (HOST, uri, slugify(slug)))

    if r.status_code != 200:
        print r.text
        return None
    #     r = requests.post('%s/category' % HOST,
    #                       json={'title': category, 'slug': slugify(category)})

    return r.json()


def getOrMakePerformer(performer):
    objectId = findObjectID('performer', performer)

    if not objectId:
        print "Making new requests for %s" % performer
        r = requests.post('%s/performer' % HOST,
                          json={'name': performer, 'slug': slugify(performer)})
        if r.status_code != 200:
            print r.text
            exit(0)

        return r.json()
    return objectId


def movie(path=None):

    data = openDataFile("data/movie.json")

    # we need to get all of the

    rank = 1
    for movie in data:
        # categories = []
        # for category in movie['Category']:

        #     # Add on first run
        #     # categories[category] = True
        #     # for category, _ in categories.iteritems():
        #     # r = requests.post('%s/category' % HOST,
        #     #   json={'title': category, 'slug': slugify(category)})
        #     print "category %s" % category
        #     categories.append(findObjectID('category', category)['_id'])

        # print "found %d categories" % len(categories)
        # performers = []
        # for performer in movie['Stars']:
        #     performers.append(getOrMakePerformer(performer)['_id'])

        # if len(movie['Director']) > 0:
        #     director = getOrMakePerformer(movie['Director'][0])
        #     director = director['_id']
        # else:
        #     director = None

        retval = {
            "category": movie['Category'],
            "information": {
                "director": movie['Director'],
                "studio": "Evil Angel",
                "performers": movie['Stars']
            },
            "playlist": [],
            "description": movie['Description'],
            "views": 1,
            "price": 1.11,
            "title": movie['Title'],
            "rank": rank,
            "reviewed": False,
            "length": length(movie['Length']),
            "upvotes": 1,
            "slug": slugify(movie['Title']),
            "downvotes": 1,
            "is_published": False
        }

        pp(json.dumps(retval))
        r = requests.post('%s/movie' % HOST,
                          data=json.dumps(retval))
        pp(r.json())
        # time.sleep(.1)
        rank += 1

# pp(data[0])


def main():

    # categories
    category(openDataFile("data/categories.json"))
    directors(openDataFile("data/performers/directors.json"))
    loadToCache("/Users/rhoop/Downloads/Evil Angel/Movie", "movie")
    movie()
    exit(0)


if __name__ == '__main__':
    main()
