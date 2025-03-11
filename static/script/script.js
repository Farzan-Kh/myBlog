document.addEventListener("DOMContentLoaded", function() {
    const articles = [
        { date: "2025-01-10", title: "WorstFit: Unveiling Hidden Transformers in Windows ANSI!" },
        { date: "2024-08-09", title: "Confusion Attacks: Exploiting Hidden Semantic Ambiguity in Apache HTTP Server!" },
        { date: "2024-06-07", title: "CVE-2024-4577 - Yet Another PHP RCE: Make PHP-CGI Argument Injection Great Again!" },
        { date: "2023-08-12", title: "從 2013 到 2023: Web Security ⼗年之進化與趨勢!" },
        { date: "2022-10-19", title: "A New Attack Surface on MS Exchange Part 4 - ProxyRelay!" },
        { date: "2022-08-18", title: "Let's Dance in the Cache - Destabilizing Hash Table on Microsoft IIS!" }
    ];

    const articleList = document.getElementById("article-list");
    articles.forEach(article => {
        let listItem = document.createElement("li");
        listItem.innerHTML = `<strong>${article.date}</strong><a href="#">${article.title}</a>`;
        articleList.appendChild(listItem);
    });
});
