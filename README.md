# GMRISA

Based on the [python project MRISA](https://github.com/vivithemage/mrisa) with added features, but done in my language of choice, [Go](https://golang.org) :D

---

GMRISA (**G**o **M**eta **R**everse **I**mage **S**earch **A**PI) is a RESTful API which takes an image URL, does a reverse Google image search, and returns a JSON array with the search results.

## Features

* Similar feature set to that of the original
* Caching
* Returns knowledge base Information if found

## API

### Perform a Reverse Image Search

Performs a reverse image search using the supplied image URL as input.

### URL

- POST to *<http://localhost:9991/search>*

### Arguments

- *__image_url__* - A URL to an image to use for the search input.
- *__image_md5__* - A MD5 of the image in question, truth be told it's just a caching key so could be anything

### Request Example:


#### CURL

```shell
curl -X POST http://localhost:9991/search \
	-H "Content-Type: application/json" \
	-H "api-key: 1234" \
	-d '{
		"image_url": "https://upload.wikimedia.org/wikipedia/en/e/e4/SydneyOperaHouse20182.jpg",
		"image_md5": "e44bda543a313bb72ea2fff6bd45b9f5"
		}'
```

#### XMLHttpRequest

```javascript

	var xhr = new XMLHttpRequest();
	xhr.open('POST',"http://localhost:9991/search");

	xhr.setRequestHeader("Content-Type","application/json");
	xhr.setRequestHeader("api-key","1234");
	xhr.onreadystatechange = () => {
		console.log(xhr.responseText);
	};;

	xhr.send(JSON.stringify({
		"image_url": "https://upload.wikimedia.org/wikipedia/en/e/e4/SydneyOperaHouse20182.jpg",
		"image_md5": "e44bda543a313bb72ea2fff6bd45b9f5"
	}));

```

### Response

There's a minor variation on the original, while the original just output all the information in seperate arrays, this returns them in results

<details>

<summary> Expand to view </summary>

<br>

```json
{
	"Entries": [
		{
			"Link": "https://www.sydneyoperahouse.com/",
			"Title": "Sydney Opera House",
			"Description": "With over 40 shows a week at the Sydney Opera House there's something for everyone. Events, tours, kids activities, food and drink - find out what's on and get ..."
		},
		{
			"Link": "https://en.wikipedia.org/wiki/Sydney_Opera_House",
			"Title": "Sydney Opera House - Wikipedia",
			"Description": "The Sydney Opera House is a multi-venue performing arts centre at Sydney Harbour in Sydney, New South Wales, Australia. It is one of the 20th century's most ..."
		},
		{
			"Link": "https://moovitapp.com/index/en-gb/public_transportation-Sydney_Opera_House-Sydney-site_7166357-442",
			"Title": "How to get to Sydney Opera House in Sydney by Train or Bus | Moovit",
			"Description": "410 × 274 - Moovit gives you the best routes to Sydney Opera House using public transport. Free step-by-step journey directions and updated timetables for Train or Bus in ...410 × 274 - "
		},
		{
			"Link": "https://www.ticketswap.com/location/sydney-opera-house/91277",
			"Title": "Sydney Opera House – Buy and sell tickets – TicketSwap",
			"Description": "2047 × 1369 - Hannah Gadsby: Douglas. Fri, December 20 • Sydney Opera House, AU · Hannah Gadsby: Douglas. Wed, December 18 • Sydney Opera House, AU ...2047 × 1369 - "
		},
		{
			"Link": "https://www.wikiwand.com/en/Sydney_Opera_House",
			"Title": "Sydney Opera House - Wikiwand",
			"Description": "540 × 361 - The Sydney Opera House is a multi-venue performing arts centre at Sydney Harbour in Sydney, New South Wales, Australia. It is one of the 20th century's most ...540 × 361 - "
		},
		{
			"Link": "https://moovitapp.com/index/en-gb/public_transportation-Sydney_Opera_House-Sydney-site_62635439-442",
			"Title": "How to get to Sydney Opera House in Sydney by Bus, Train or Ferry ...",
			"Description": "410 × 274 - Moovit gives you the best routes to Sydney Opera House using public transport. Free step-by-step journey directions and updated timetables for Bus, Train or ...410 × 274 - "
		},
		{
			"Link": "https://www.ticketswap.uk/location/sydney-opera-house/91277",
			"Title": "Sydney Opera House – Buy and sell tickets – TicketSwap",
			"Description": "2047 × 1369 - Megan Mullally and her band Nancy And Beth. Sun, June 16 • Sydney Opera House, AU ... À Ố Làng Ph . Fri, June 14 • Sydney Opera House, AU ...2047 × 1369 - "
		}
	],
	"Similar": [
		"https://www.gettyimages.com/photos/sydney-opera-house",
		"https://www.gettyimages.com/detail/photo/sydney-opera-house-in-the-sun-royalty-free-image/869714270",
		"https://www.gettyimages.com/photos/sydney-opera-house",
		"https://www.createdigital.org.au/australian-new-zealand-engineering-projects-world/",
		"https://sydneyexpert.com/photograph-the-sydney-opera-house/",
		"https://www.ytravelblog.com/what-to-do-in-sydney/",
		"https://www.gettyimages.com/detail/video/wide-shot-time-lapse-boats-on-harbor-around-sydney-stock-video-footage/753-81",
		"http://content.time.com/time/world/article/0,8599,2097247,00.html",
		"https://www.stayatbase.com/sydney/must-dos-in-sydney/",
		"https://www.shutterstock.com/image-photo/sydney-july-8-opera-house-view-111266186"
	],
	"BestGuess": "sydney opera house",
	"KnowledgeBase": {
		"Title": "Sydney Opera House",
		"Subtitle": "Performing arts centre in the City of Sydney, New South Wales",
		"Description": "The Sydney Opera House is a multi-venue performing arts centre at Sydney Harbour in Sydney, New South Wales, Australia. It is one of the 20th century's most famous and distinctive buildings. Wikipedia ",
		"Modules": {
			"Architect": [
				"Jørn Utzon"
			],
			"Architecture firm": [
				"Arup Group"
			],
			"Construction started": [
				"2 March 1959"
			],
			"Did you know": [
				"Paul Robeson was the first to perform in the Opera House."
			],
			"Founded": [
				"1973"
			]
		}
	}
}
```

</details>
