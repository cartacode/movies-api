import string
import json
from slugify import slugify
from requests import get, post
from bs4 import BeautifulSoup
from contextlib import closing
from requests.exceptions import RequestException
from datetime import datetime
import re
import pprint
pp = pprint.pprint


def openDataFile(path):

    data = None

    with open(path,) as fileh:
        data = json.load(fileh)
    return data


def myconverter(o):
    if isinstance(o, datetime):
        return o.utcnow


def unique(list1):

    # intilize a null list

    unique_list = []

    # traverse for all elements

    for x in list1:

        # check if exists in unique_list or not

        if x not in unique_list:

            unique_list.append(x)

    # print list

    return unique_list


def simple_get(url):
    """
    Attempts to get the content at `url` by making an HTTP GET request.
    If the content-type of response is some kind of HTML/XML, return the
    text content, otherwise return None.
    """
    try:
        with closing(get(url, stream=True)) as resp:
            if is_good_response(resp):
                return resp.content
            else:
                return None

    except RequestException as e:
        log_error('Error during requests to {0} : {1}'.format(url, str(e)))
        return None


def is_good_response(resp):
    """
    Returns True if the response seems to be HTML, False otherwise.
    """
    content_type = resp.headers['Content-Type'].lower()
    return (resp.status_code == 200 and
            content_type is not None and
            content_type.find('html') > -1)


def log_error(e):
    """
    It is always a good idea to log errors.
    This function just prints them, but you can
    make it do anything.
    """
    print(e)
    exit(0)


def main():
    data = openDataFile("data/movie.json")

    # we need to get all of the

    stars = []
    for movie in data:
        for star in movie['Stars']:
            stars.append(star)

    for star in unique(stars):
        raw_html = simple_get(
            "https://www.pornhub.com/pornstar/%s" % slugify(star))
        if not raw_html:
            print "Can't find info for [%s]" % star
            continue
        print "Found info for [%s]" % star
        soup = BeautifulSoup(raw_html, "lxml")
        info = soup.findAll("div", {"class": "infoPiece"})

        catch = {"enabled": False, "key": ""}
        profile = {
            "slug": slugify(star),
            "name": star,
            "birthdate": None,
            "birthplace": None,
            "size": {},
            "traits": {}

        }

        for div in info:
            # pp(div)
            data = string.lstrip(div.text)
            k, v = data.split(":\n")
            # print k
            # print "[%s]" % data
            if k == "Gender":
                profile["gender"] = v
            elif k == "Birthday":
                profile["birthdate"] = datetime.strptime(
                    string.strip(v), '%b %d, %Y')
            elif k == "Star Sign":
                profile["traits"]["sign"] = v
            elif k == "Measurements":
                profile["size"]["bust"] = v
            elif k == "Height":
                regex = r"\(\d+"
                match = re.search(regex, v)
                # for match in matches:

                profile["size"]["height"] = int(0.393701 *
                                                float(match.group(0).split('(')[1]))
            elif k == "Weight":
                regex = r"\d+"
                match = re.search(regex, v)
                # for match in matches:
                profile["size"]["weight"] = int(match.group(0))
            elif k == "Ethnicity":
                profile["traits"]["ethnicity"] = v
            elif k == "Hair Color":
                profile["traits"]["hair_color"] = v
            elif k == "Tattoos":
                profile["traits"]["tattoos"] = bool(v)
            elif k == "Piercings":
                profile["traits"]["piercings"] = bool(v)
            elif k == "City and Country":
                profile["birthplace"] = v
            elif k == "Born":
                profile["birthdate"] = datetime.strptime(
                    string.strip(v), '%Y-%m-%d')

        pp(profile)
        r = post('http://127.0.0.1:3000/performer',
                 data=json.dumps(profile, default=myconverter))

        print r.text
        # exit(0)
        # for span in div.find_all('Height:', recursive=True):
        #     for child in span.children:
        #         print child,
        # print stars
    exit(0)


if __name__ == '__main__':
    main()

#

