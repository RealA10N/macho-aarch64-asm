"""
This script generated the CSV AArch64BaseInstructions.csv file
"""

import csv
import typing

from bs4 import BeautifulSoup  # pip install beautifulsoup4


def parse_html(html: str) -> list[tuple[str, str, str]]:
    """parse the HTML and return a tuple of instructions in the format
    (name, description, url)."""

    soup = BeautifulSoup(html, "lxml")
    container = soup.find(attrs={"class": "section-wrapper"})
    divs = container.find_all("div")

    items = list()
    for div in divs:
        name, _, description = div.text.partition(":")
        url = div.find("a")["href"]
        items.append((name, description, url))

    return items


def save_data(data: list[tuple[str, str, str]], file: typing.TextIO) -> None:
    writer = csv.writer(file)
    writer.writerow(["Instruction Name", "Description", "Url"])
    for row in data:
        writer.writerow(row)


def main() -> None:
    # The HTML at the following URL after Javascript rendering
    # https://developer.arm.com/documentation/ddi0596/2020-12/Base-Instructions?lang=en
    input_filepath = "Arm A64 Instruction Set Architecture.html"

    # Where to save the resulting CSV file
    output_filepath = "AArch64BaseInstructions.csv"

    with open(input_filepath, "r", encoding="utf8") as file:
        html = file.read()

    data = parse_html(html)

    with open(output_filepath, "w", encoding="utf8") as file:
        save_data(data, file)
