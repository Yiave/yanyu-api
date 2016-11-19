# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# http://doc.scrapy.org/en/latest/topics/items.html

from scrapy.item import Item, Field


class University(Item):
    school = Field()
    location = Field()
    subjection = Field()
    attribute = Field()
    is_graduate_school = Field()
    self_decision = Field()


#import scrapy


#    class SchoolDataItem(scrapy.Item):
        # define the fields for your item here like:
        # name = scrapy.Field()

 #       pass
