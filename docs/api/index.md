---
layout: default
title: API Specification
nav_order: 4
permalink: /api/
classes: redoc-wide
---

<div class="redoc-wide">
  <div id="redoc-container" style="min-height: 85vh;"></div>
</div>

<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js" defer></script>
<script>
  window.addEventListener('DOMContentLoaded', function () {
    Redoc.init(
      '{{ site.baseurl }}/api/openapi.yaml',
      {
        scrollYOffset: 60,
        hideDownloadButton: false,
        expandResponses: "200,201",
        theme: {
          colors: {
            primary: { main: "#2b6cb0" }
          },
          typography: {
            fontSize: "16px",
            lineHeight: "1.6"
          }
        }
      },
      document.getElementById('redoc-container')
    );
  });
</script>
