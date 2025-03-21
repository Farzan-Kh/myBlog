# Personal Blog

This is a personal blog project built with Go, designed to showcase posts, articles, and other content. The blog uses **Strapi**, a headless CMS, for content management and integrates various features such as RSS feeds, a newsletter subscription, and a responsive design.

## Features

- **Dynamic Content Management**: Content is fetched from Strapi, allowing easy updates and management.
- **Responsive Design**: The blog is fully responsive and works seamlessly across devices.
- **RSS Feed**: Automatically generates an RSS feed for posts.
- **Newsletter Subscription**: Allows users to subscribe to a newsletter.
- **Search Functionality**: Users can search for posts using a search bar.
- **Table of Contents**: Automatically generates a table of contents for posts.
- **Static Assets**: Includes static assets such as images, styles, and scripts.

## Project Structure

```
.
├── cmd/
│   └── myBlog/
│       └── main.go          # Entry point of the application
├── dummyData/
│   └── dummyPosts.json      # Dummy data for local development
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── handlers/
│   │   ├── about.go         # Handler for the About page
│   │   ├── home.go          # Handler for the Home page
│   │   ├── posts.go         # Handler for Posts
│   │   ├── rss.go           # Handler for RSS feed
│   │   └── search.go        # Handler for Search functionality
│   ├── models/
│   │   └── posts.go         # Data model for blog posts
│   └── utils/
│       ├── logger.go        # Logging utilities
│       ├── template.go      # Template rendering utilities
│       └── timer.go         # Timer utility for periodic tasks
├── static/
│   ├── img/                 # Static images
│   ├── script/              # JavaScript files
│   │   ├── main.js          # Main script for the blog
│   │   └── post.js          # Script for post-specific functionality
│   └── style/
│       └── styles.css       # CSS styles
├── templates/
│   ├── about.html           # Template for the About page
│   ├── home.html            # Template for the Home page
│   └── post.html            # Template for individual posts
├── logs/
│   └── logs.txt             # Log file for application events
├── .env                     # Environment variables
├── go.mod                   # Go module file
└── go.sum                   # Go dependencies checksum
```

## Prerequisites

- **Go**: Version 1.23.2 or higher.
- **Strapi**: A running instance of Strapi for content management.
- **Node.js**: For managing Strapi (if running locally).
- **Environment Variables**: Configure the `.env` file with the following variables:
  - `API_ADDRESS`: The base URL of the Strapi instance.
  - `API_KEY`: The API key for accessing Strapi.
  - `DEBUG`: Set to `TRUE` for development mode.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/personal-blog.git
   cd personal-blog
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the .env file:
   ```bash
   cp .env.example .env
   # Update the .env file with your Strapi API details
   ```

4. Run the application:
   ```bash
   go run cmd/myBlog/main.go
   ```

## Usage

- **Home Page**: Accessible at home, displays a list of blog posts.
- **Post Page**: Accessible at `/posts/{UUID}`, displays individual posts.
- **About Page**: Accessible at `/about`, displays information about the blog owner.
- **RSS Feed**: Accessible at `/rss.xml`, provides an RSS feed of the blog posts.
- **Newsletter Registration**: Accessible via the home page modal.

## Development

### Fetching Posts
The application fetches posts from Strapi using the `FetchPosts` function. In development mode (`DEBUG=TRUE`), it uses dummy data from `dummyPosts.json`.

### Logging
Logs are stored in logs.txt. The application uses `utils.InfoLogger` and `utils.ErrorLogger` for logging.

### Templates
HTML templates are located in the `templates` directory. They use Go's `html/template` package for rendering.

## Deployment

1. Build the application:
   ```bash
   go build -o blog cmd/myBlog/main.go
   ```

2. Deploy the binary and static files to your server.

3. Ensure the .env file is configured correctly on the server.

4. Start the application:
   ```bash
   ./blog
   ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- [Strapi](https://strapi.io/) for content management.
- [Go](https://golang.org/) for powering the backend.
- [Bootstrap](https://getbootstrap.com/) for responsive design.
- [Font Awesome](https://fontawesome.com/) for icons.
