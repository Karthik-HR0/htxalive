<h1 align="center">htxalive</h1>

> **htxalive** - Your go-to tool for HTTP status checks, title extraction, and IP resolution.

`htxalive` is a fast and efficient command-line utility written in Go, designed to streamline web reconnaissance. Whether you're analyzing URLs for availability, extracting page titles, or resolving IP addresses, `htxalive` makes the process seamless and efficient.

---

### Features

- **Fetch HTTP Status Codes**: Instantly check the availability of URLs.
- **Extract Page Titles**: Quickly gather the titles of web pages.
- **Resolve IP Addresses**: Map domains to their IPs.
- **Concurrency and Threading**: Blazing-fast processing of URLs.
- **Output to File**: Save results for future reference.
- **Minimal Mode**: Keep things clean with silent output.

---

### Installation

Install `htxalive` directly from the source:

```bash
go install github.com/Karthik-HR0/htxalive@latest
```


---

Arguments


---

Usage Examples

1. Check HTTP status codes:

```bash 
echo "https://example.com" | htxalive -sc
```

2. Display page titles and IPs:

```bash
echo "https://example.com" | htxalive -tl -ip
```

3. Save output to a file with concurrency:

```bash
cat urls.txt | htxalive -c 20 -o output.txt
```

4. Filter results for live URLs:

```bash
 cat urls.txt | htxalive -tl | grep "200" > live_urls.txt
```

5. Save to file and count results:

```bash
cat urls.txt | htxalive | tee results.txt | wc -l
```


---

Troubleshooting

Empty Results?

Ensure URLs start with http:// or https://.

Check your internet connection.

Increase concurrency and threading for large input files.


Output Issues?

Verify file permissions for the -o argument.




---

<p align="center">
Built with ❤️ by <a href="https://github.com/Karthik-HR0">@Karthik-HR0</a>
</p>
