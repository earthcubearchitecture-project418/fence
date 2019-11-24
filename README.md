# Fence

## About

Fence is a tool designed to allow people to evaluate structured data graphs against various validations (such at the Tangram SHACL service), tools and interface components.

It also provides links to the Google Structured Data Testing Tool and the JSON-LD Playground.  

A set of web components are also loaded and attempt to parse the JSON-LD data graph to testing approaches to mapping, citation generation and other views into the data graph.   This is done to provide examples of how providers can further leverage their data graphs in the generation of data set landing pages.

More information will be coming soon on how to use Fence.  

You can visit it for now at http://fence.gleaner.io/fence .

To try it testing a live resource use:


https://fence.gleaner.io/fence?url=http://opencoredata.org/doc/dataset/b8d7bd1b-ef3b-4b08-a327-e28e1420adf0 

or 

https://fence.gleaner.io/fence?url=https://www.bco-dmo.org/dataset/722472 

as examples.  

## Fence pull

### About

This is a function that allows you to pull the JSON-LD from a given page.  This
is a utility function and documentation is coming.  

## Headless 

### About

We will add headless procesing to this soon and just switch over to using this
fully.  It will complicate deployment a bit as I will need to be able to deploy
the chromedp container to be the headless server.   

```
docker pull chromedp/headless-shell
```



