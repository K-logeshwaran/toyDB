import requests
import time
import sys 
dummy_data = [
    {
        "age": 25,
        "dept": "45th",
        "name": "Alice",
        "rem": "New hire"
    },
    {
        "age": 38,
        "dept": "12th",
        "name": "Bob",
        "rem": "Experienced"
    },
    {
        "age": 42,
        "dept": "33rd",
        "name": "Cynthia",
        "rem": "Team lead"
    },
    {
        "age": 29,
        "dept": "7th",
        "name": "David",
        "rem": "Skilled developer"
    },
    {
        "age": 55,
        "dept": "22nd",
        "name": "Eleanor",
        "rem": "Seasoned manager"
    },
    {
        "age": 31,
        "dept": "18th",
        "name": "Frank",
        "rem": "Creative designer"
    },
    {
        "age": 40,
        "dept": "50th",
        "name": "Grace",
        "rem": "Marketing guru"
    },
    {
        "age": 27,
        "dept": "9th",
        "name": "Henry",
        "rem": "Junior developer"
    },
    {
        "age": 48,
        "dept": "28th",
        "name": "Isabella",
        "rem": "Experienced engineer"
    },
    {
        "age": 33,
        "dept": "15th",
        "name": "Jack",
        "rem": "Tech enthusiast"
    },
    {
        "age": 22,
        "dept": "4th",
        "name": "Katherine",
        "rem": "Fresh recruit"
    },
    {
        "age": 37,
        "dept": "39th",
        "name": "Liam",
        "rem": "Senior analyst"
    },
    {
        "age": 29,
        "dept": "21st",
        "name": "Mia",
        "rem": "Upcoming talent"
    },
    {
        "age": 44,
        "dept": "8th",
        "name": "Noah",
        "rem": "Seasoned developer"
    },
    {
        "age": 31,
        "dept": "31st",
        "name": "Olivia",
        "rem": "Creative thinker"
    },
    {
        "age": 26,
        "dept": "6th",
        "name": "Patrick",
        "rem": "Junior designer"
    },
    {
        "age": 35,
        "dept": "13th",
        "name": "Quinn",
        "rem": "Versatile employee"
    },
    {
        "age": 42,
        "dept": "17th",
        "name": "Rachel",
        "rem": "Experienced manager"
    },
    {
        "age": 28,
        "dept": "20th",
        "name": "Samuel",
        "rem": "Tech enthusiast"
    },
    {
        "age": 51,
        "dept": "25th",
        "name": "Tracy",
        "rem": "Skilled developer"
    },
    {
        "age": 32,
        "dept": "32nd",
        "name": "Ulysses",
        "rem": "Creative designer"
    },
    {
        "age": 23,
        "dept": "2nd",
        "name": "Victoria",
        "rem": "Fresh recruit"
    },
    {
        "age": 29,
        "dept": "11th",
        "name": "William",
        "rem": "Eager learner"
    },
    {
        "age": 46,
        "dept": "27th",
        "name": "Xavier",
        "rem": "Experienced developer"
    },
    {
        "age": 34,
        "dept": "19th",
        "name": "Yasmine",
        "rem": "Innovative thinker"
    },
    {
        "age": 41,
        "dept": "36th",
        "name": "Zachary",
        "rem": "Senior engineer"
    },
    {
        "age": 30,
        "dept": "42nd",
        "name": "Aria",
        "rem": "Talented designer"
    },
    {
        "age": 28,
        "dept": "10th",
        "name": "Benjamin",
        "rem": "Creative mind"
    },
    {
        "age": 39,
        "dept": "29th",
        "name": "Chloe",
        "rem": "Proven expertise"
    },
    {
        "age": 36,
        "dept": "48th",
        "name": "Daniel",
        "rem": "Versatile professional"
    },
    {
        "age": 50,
        "dept": "16th",
        "name": "Emily",
        "rem": "Experienced leader"
    },
    {
        "age": 27,
        "dept": "14th",
        "name": "Fiona",
        "rem": "Tech enthusiast"
    },
    {
        "age": 33,
        "dept": "23rd",
        "name": "George",
        "rem": "Innovator"
    },
    # ... continue with more objects
]
url = sys.argv[1]

# Sending the POST request
for i in dummy_data:

    response = requests.post(url, json=i)

    # Handling the response
    if response.status_code == 200:
        print("POST request was successful!")
        print("Response:", response.content)
    else:
        print("POST request failed with status code:", response.status_code)
        #print("Response:", response.text)
    time.sleep(2)

print(f"no of data added {len(dummy_data)}")