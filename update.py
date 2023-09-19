#!/bin/python3

import json
import os
import requests

entities_link = "https://raw.githubusercontent.com/disconnectme/disconnect-tracking-protection/master/entities.json"
services_link = "https://raw.githubusercontent.com/disconnectme/disconnect-tracking-protection/master/services.json"

def retrieve_contents(link):
    response = requests.get(link)
    return response.content

def generate_entities_files():
    content = retrieve_contents(entities_link)
    entities = json.loads(content)["entities"]

    domains = set()
    for entity in entities.values():
        for domain in entity["resources"]:
            domain = domain.split('/')[0].strip()  # strip out path from domain
            domains.add(domain)

    with open("entities.txt", "w") as f:
        for domain in sorted(domains):
            print(domain, file=f)

def generate_services_files():
    content = retrieve_contents(services_link)
    categories = json.loads(content)["categories"]

    category_domains = {}
    for category, services in categories.items():
        for service in services:
            for domains in service.values():
                for domain_list in domains.values():
                    if isinstance(domain_list, list):
                        domain_list = [domain.split('/')[0] for domain in domain_list]  # strip out path from domain
                        category_domains.setdefault(category, set()).update(domain_list)

    with open("services.txt", "w") as f:
        for category, domains in sorted(category_domains.items()):
            for domain in sorted(domains):
                print(domain, file=f)

    for category, domains in sorted(category_domains.items()):
        with open(f"services_{category}.txt", "w") as f_category:
            for domain in sorted(domains):
                print(domain, file=f_category)

generate_entities_files()
generate_services_files()
