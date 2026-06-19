(function () {
  'use strict';

  // ── Slug helpers ─────────────────────────────────────────────────────
  // Slugs look like "north_carolina-usa" or "paris-france"
  // Split on the LAST hyphen so hyphenated city names survive.

  function titlify(str) {
    return str.split('_').map(function (w) {
      return w.charAt(0).toUpperCase() + w.slice(1);
    }).join('\u00a0');   // non-breaking spaces inside a city name
  }

  function parseSlug(slug) {
    var dash = slug.lastIndexOf('-');
    if (dash === -1) return { city: titlify(slug), country: '' };
    return {
      city:    titlify(slug.slice(0, dash)),
      country: titlify(slug.slice(dash + 1))
    };
  }

  // Populate city / country spans
  var locationCards = document.querySelectorAll('.location-card[data-slug]');
  locationCards.forEach(function (card) {
    var p = parseSlug(card.dataset.slug);
    var cityEl    = card.querySelector('.loc-city');
    var countryEl = card.querySelector('.loc-country');
    if (cityEl)    cityEl.textContent    = p.city;
    if (countryEl) countryEl.textContent = p.country;
    // Hide the separator if there is no country
    if (!p.country) {
      var sep = card.querySelector('.loc-sep');
      if (sep) sep.hidden = true;
    }
  });

  // ── Date helpers ─────────────────────────────────────────────────────
  // Dates arrive as "dd-mm-yyyy" (past) or "*dd-mm-yyyy" (upcoming).

  var MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

  function formatDate(raw) {
    var upcoming = raw.charAt(0) === '*';
    var s        = upcoming ? raw.slice(1) : raw;
    var parts    = s.split('-');

    if (parts.length !== 3) return { text: s, upcoming: upcoming };

    var day   = parseInt(parts[0], 10);
    var month = parseInt(parts[1], 10) - 1;   // 0-based
    var year  = parts[2];
    var text  = day + '\u00a0' + (MONTHS[month] || parts[1]) + '\u00a0' + year;

    return { text: text, upcoming: upcoming };
  }

  var dateItems = document.querySelectorAll('.date-item[data-raw]');
  dateItems.forEach(function (item) {
    var f = formatDate(item.dataset.raw);
    // Prefix upcoming shows with an up-arrow
    item.textContent = f.upcoming ? '\u2191\u2009' + f.text : f.text;
    if (f.upcoming) item.classList.add('date-item--upcoming');
  });

})();
