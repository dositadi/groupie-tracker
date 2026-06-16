(function () {
  var grid = document.getElementById('artist-grid');
  if (!grid) return;

  var cards = Array.prototype.slice.call(grid.querySelectorAll('.card'));
  var total = cards.length;
  if (total === 0) return;

  var searchInput   = document.getElementById('search-input');
  var eraSelect     = document.getElementById('era-filter');
  var countrySelect = document.getElementById('country-filter');
  var sortSelect    = document.getElementById('sort-select');
  var clearBtn      = document.getElementById('clear-filters');
  var resultLine    = document.getElementById('result-count');
  var noResults     = document.getElementById('no-results');

  // Build era + country options from what's actually in the data.
  var decades   = new Set();
  var countries = new Map();

  cards.forEach(function (card) {
    var year = parseInt(card.dataset.creation, 10);
    if (!isNaN(year)) decades.add(Math.floor(year / 10) * 10);

    card.dataset.locations.split('|').filter(Boolean).forEach(function (loc) {
      var idx = loc.indexOf('-');
      if (idx === -1) return;
      var slug = loc.slice(idx + 1);
      if (countries.has(slug)) return;
      var label = slug.split('_').map(function (w) {
        return w.charAt(0).toUpperCase() + w.slice(1);
      }).join(' ');
      countries.set(slug, label);
    });
  });

  Array.from(decades).sort(function (a, b) { return a - b; }).forEach(function (d) {
    var opt = document.createElement('option');
    opt.value = d;
    opt.textContent = d + 's';
    eraSelect.appendChild(opt);
  });

  Array.from(countries.entries())
    .sort(function (a, b) { return a[1].localeCompare(b[1]); })
    .forEach(function (entry) {
      var opt = document.createElement('option');
      opt.value = entry[0];
      opt.textContent = entry[1];
      countrySelect.appendChild(opt);
    });

  function applyFilters() {
    var term    = searchInput.value.trim().toLowerCase();
    var era     = eraSelect.value;
    var country = countrySelect.value;
    var visible = 0;

    cards.forEach(function (card) {
      var haystack      = (card.dataset.name + ' ' + card.dataset.members).toLowerCase();
      var matchesSearch = term === '' || haystack.indexOf(term) !== -1;

      var matchesEra = era === '';
      if (!matchesEra) {
        var year = parseInt(card.dataset.creation, 10);
        matchesEra = Math.floor(year / 10) * 10 === parseInt(era, 10);
      }

      var matchesCountry = country === '';
      if (!matchesCountry) {
        matchesCountry = card.dataset.locations.split('|').filter(Boolean).some(function (loc) {
          return loc.slice(loc.indexOf('-') + 1) === country;
        });
      }

      var show = matchesSearch && matchesEra && matchesCountry;
      card.classList.toggle('is-hidden', !show);
      if (show) visible++;
    });

    resultLine.textContent = 'Showing ' + visible + ' of ' + total + ' artists';
    if (noResults) noResults.hidden = visible !== 0;

    var active = term !== '' || era !== '' || country !== '';
    clearBtn.classList.toggle('is-active', active);
  }

  function applySort() {
    var value  = sortSelect.value;
    var sorted = cards.slice();

    if (value === 'name-asc') {
      sorted.sort(function (a, b) { return a.dataset.name.localeCompare(b.dataset.name); });
    } else if (value === 'name-desc') {
      sorted.sort(function (a, b) { return b.dataset.name.localeCompare(a.dataset.name); });
    } else if (value === 'formed-new') {
      sorted.sort(function (a, b) { return parseInt(b.dataset.creation, 10) - parseInt(a.dataset.creation, 10); });
    } else if (value === 'formed-old') {
      sorted.sort(function (a, b) { return parseInt(a.dataset.creation, 10) - parseInt(b.dataset.creation, 10); });
    }

    sorted.forEach(function (card) { grid.appendChild(card); });
  }

  searchInput.addEventListener('input', applyFilters);
  eraSelect.addEventListener('change', applyFilters);
  countrySelect.addEventListener('change', applyFilters);
  sortSelect.addEventListener('change', applySort);

  clearBtn.addEventListener('click', function () {
    searchInput.value = '';
    eraSelect.value   = '';
    countrySelect.value = '';
    applyFilters();
  });

  // Hamburger toggle
  var filterToggle = document.getElementById('filter-toggle');
  var filterPanel  = document.getElementById('filter-panel');

  if (filterToggle && filterPanel) {
    filterToggle.addEventListener('click', function () {
      var open = filterPanel.classList.toggle('is-open');
      filterToggle.setAttribute('aria-expanded', String(open));
    });
  }

  applyFilters();
})();