"""
* You are building instagram reels where users can upload and delete reels/videos.

* Each reel is a string of digits, where the ith digit of the string represents the 
content of the reel at minute i. For example, the first digit represents the content at 
minute 0 in the reel, the second digit represents the content at minute 1 in the reel, 
and so on. Viewers of reels can also like and dislike reels.

* Internally, the platform keeps track of the number of views, likes, and dislikes on each reel.

* Implement functionality for a ReelPlatform class including but not limited to:
  * upload(reel: str) → return the ID of the reel
  * remove(video_id) → return bool if removal works
  * watch(video_id, start, end) → return string (chunk of reel from start to end)
    - This should increment the view count on a valid watch
  * like(video_id) → None
  * dislike(video_id) → None
  * get_likes_and_dislikes(video_id) → return dict of likes and dislikes
  * get_views(video_id) → return int - the number of views
"""

from uuid import uuid4

class ReelPlatform() :
    def __init__(self):
        self.reel = {}

    def upload(self, reelContent ):
        id = uuid4()
        self.reel[id] = {reelContent: [0,0,0]} 
        return id

    def get_views(self, id):   
        reel =  self.reel[id]
        data = reel.values
        return data.values

    # def delete(self):
    # def like(self):   
    # def dislike(self):  
    # def get_likes_and_dislikes(self):  
          

rp = ReelPlatform()
reelID = rp.upload("123")
views = rp.get_views(reelID)
print("This is the reel id that was uploaded : ", reelID, views)


