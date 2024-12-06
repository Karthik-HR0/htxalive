<h1 align="center">htxalive</h1>

> **htxalive**  - A reliable tool to check the "life" status of HTTP transactions ,status checks, title extraction, and IP resolution.

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

Uses


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

<h3>Arguments</h3>

<table>
  <thead>
    <tr>
      <th>Argument</th>
      <th>Description</th>
      <th>Default</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>-o</code></td>
      <td>Save output URLs to a file</td>
      <td>None</td>
    </tr>
    <tr>
      <td><code>-s</code></td>
      <td>Silent mode (minimal output)</td>
      <td><code>false</code></td>
    </tr>
    <tr>
      <td><code>-c</code></td>
      <td>Concurrency level</td>
      <td><code>10</code></td>
    </tr>
    <tr>
      <td><code>-t</code></td>
      <td>Threading level</td>
      <td><code>3</code></td>
    </tr>
    <tr>
      <td><code>-sc</code></td>
      <td>Show HTTP status codes</td>
      <td><code>false</code></td>
    </tr>
    <tr>
      <td><code>-tl</code></td>
      <td>Show page titles</td>
      <td><code>false</code></td>
    </tr>
    <tr>
      <td><code>-ip</code></td>
      <td>Show resolved IP addresses</td>
      <td><code>false</code></td>
    </tr>
  </tbody>
</table>

---
```bash

.__     __ ____  ___      .__  .__
|  |___/  |\   \/  /____  |  | |__|__  __ ____
|  |  \   __\     /\__  \ |  | |  \  \/ // __ \
|   Y  \  | /     \ / __ \|  |_|  |\   /\  ___/
|___|  /__|/___/\  (____  /____/__| \_/  \___  >
     \/          \_/    \/                   \/

                           with <3 by @Karthik-HR0


[INF] Current htxalive version v1.0.1 (version check failed)
Usage of ./htxalive:
  -c int
        Concurrency level (default 10)
  -ip
        Show IP address
  -o string
        Collect only output URLs
  -s    Silent mode
  -sc
        Show HTTP status codes
  -t int
        Threading level (default 3)
  -tl
        Show page titles
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
