# Asio

# Roadmap 

MF - Major feature
mf - Minor feature

- [ ] Status Pages - v0.1
  - [ ] Connect to NoSQL database (preferably MongoDB) - MF
  - [ ] Create models for single domains. - MF
  - [ ] Integrate an asynchronous task queue. - MF
  - [ ] Ping to single domains (ports 80-443). (As asynchronous task called `ping_domain`) - MF
  - [ ] Calculate response time. (Within `ping_domain`) - MF
  - [ ] Get the HTML content and save URLs. - mf
  - [ ] Make basic implementation with https://github.com/PuerkitoBio/gocrawl. - mf
  - [ ] Create events model, for records of ping. - MF
  - [ ] Create REST interface (Define framework* - Options: Gorilla, Martini, Gin, Revel) - MF
  - [ ] Create endpoints for manage and get info of single domain. - MF
  - [ ] Create endpoints for manage events (events are the ping result to a single domain). - mf

## Schedule

- May 5, 2015 - Asio 0.1 alpha.
- May 15, 2015 - Asio 0.1 beta.
- May 25, 2015 - Asio 0.1 final release.

## Alpha

Branch for the alpha deadline is `stable/0.1` (https://github.com/ClowlHQ/asio/tree/stable/0.1). Freezing major features if were completed, another features will assing to the next version (milestone).

## Final release

Merge `stable/0.1` to merge.
