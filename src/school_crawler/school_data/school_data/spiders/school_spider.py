# -*- coding=UTF-8 -*-

import scrapy

from school_data.items import University
class SchoolSpiser(scrapy.Spider):
    name = "school"
    allowed_domains = ["chsi.com.cn"]
    start_urls = [
        "http://yz.chsi.com.cn/sch/?start=0"
    ]


    def parse(self, response):
        res = response.xpath('//tbody')

        items = []
        for tr in res.xpath('tr'):
            school = tr.xpath('td/a/text()').extract()[0]#.lstrip(' \n')#.rstrip(' ').replace('\n', ' ') # 删除左右的的空格和换行，并用空格替换最右边的换行
            location = tr.xpath('td/text()').extract()[1]
            subjection = tr.xpath('td/text()').extract()[2]
            attribute = tr.xpath('td[4]/span/text()').extract()
            is_graduate_school = tr.xpath('td[5]/text()')
            if is_graduate_school: # list可以直接判断是否为空
                is_graduate_school = "yes"
            else:
                is_graduate_school = "no"

            self_decision = tr.xpath('td[6]/text()')
            if self_decision:
                self_decision = "yes"
            else:
                self_decision = "no"


            item = University()
            item['school'] = school.lstrip('\r\n ').rstrip('\r\n ')
            item['location'] = location
            item['subjection'] = subjection
            item['attribute'] = attribute
            item['is_graduate_school'] = is_graduate_school
            item['self_decision'] = self_decision
            items.append(item)
            #print school, location, subjection, attribute, is_graduate_school, self_decision

        next_url = response.xpath('//form/ul/li[9]').xpath('a/@href').extract()
        if next_url:
            url = next_url[0]
            print "=======>", url
            #yield self.make_requests_from_url(url)  # add url to crawl urls
            items.append(self.make_requests_from_url(url))
        return items

