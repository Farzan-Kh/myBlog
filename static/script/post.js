document.addEventListener("DOMContentLoaded", () => {
    const tocContainer = document.getElementById("toc");
    const headings = Array.from(document.querySelectorAll("h1, h2, h3, h4, h5, h6")).slice(1);
    const tocList = document.createElement("ul");
    tocList.classList += "ms-5 row align-middle"
    tocList.id = "tocList"

    let currentLevel = 1;
    let listStack = [tocList];

    headings.forEach((heading) => {
        // Generate an ID for each heading
        const id = heading.textContent.toLowerCase().trim().replace(/[^\w\s]/g, "").replace(/\s+/g, "-");
        heading.id = id;

        // Create list item with link
        const listItem = document.createElement("li");
        const link = document.createElement("a");
        link.href = `#${id}`;
        link.textContent = heading.textContent;
        listItem.appendChild(link);

        // Determine the level of the heading (e.g., <h2> = 2)
        const level = parseInt(heading.tagName[1]);

        // Adjust the list nesting based on heading level
        if (level > currentLevel) {
            const newSubList = document.createElement("ul");
            listStack[0].lastElementChild.appendChild(newSubList);
            listStack.unshift(newSubList);
        } else if (level < currentLevel) {
            while (level < currentLevel) {
                listStack.shift();
                currentLevel--;
            }
        }

        listStack[0].appendChild(listItem);
        currentLevel = level;
    });

    tocContainer.appendChild(tocList);
});