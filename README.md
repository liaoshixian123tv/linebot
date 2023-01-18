# Linebot exercise

## Introduce
linebot exercise will receive the message from the webhook and save the message to database. It also offer a restful api which can get message from database.

## Preparation
* MongoDB need to executed
* config need to complete
* ngork 

## Run the code
1. clone the code
2. go run main.go

## Struct of code
![image](https://github.com/liaoshixian123tv/linebot/blob/main/linebot_structure.png)

## Restful API
### Get message from database
api path: `http://your_ip:your_port/getmessage`

if startTime and endTime both are empty, it will return all message.

| Header        | Description   |
| ------------- |:-------------:|
| startTime     | timestamp start|
| endTime       | timestamp end |

**Response**
```
[
    {
        "messageId": "17486870935193",
        "userId": "Ua4ca938776eddf76aa17a8e000f62e97",
        "timestamp": 1673958173304,
        "sticker": {
            "packageId": "1257552",
            "stickerId": "10443982",
            "stickerResourceType": "STATIC",
            "keywords": [
                "Negative",
                "impossible",
                "Confounded",
                "Never mind",
                "Good grief",
                "over board",
                "idk"
            ]
        },
        "location": {
            "address": "",
            "latitude": 0,
            "longitude": 0
        },
        "messageType": 6
    },
    {
        "content": "Iua4rOippiI=",
        "messageId": "17486887212105",
        "userId": "Ua4ca938776eddf76aa17a8e000f62e97",
        "timestamp": 1673958361513,
        "sticker": {
            "packageId": "",
            "stickerId": "",
            "stickerResourceType": "",
            "keywords": null
        },
        "location": {
            "address": "",
            "latitude": 0,
            "longitude": 0
        },
        "messageType": 1,
        "contentStr": "\"測試\""
    }
]
```

## Future function
1. Add database transcation function
2. Available for more types of message
 * Now available __text__, __image__,__audio__,__video__,__sticker__,__location__ 
