document.addEventListener('DOMContentLoaded', (event) => {
  const bandCards = document.querySelectorAll('.band-card');
  const searchInput = document.querySelector('#search-input');
  const suggestionsContainer = document.querySelector('#suggestions-container');
  // const featuredBands = document.querySelector('.featured-bands');

  // Add hover effect to band cards
  bandCards.forEach(card => {
    card.addEventListener('mouseenter', () => {
      card.style.transform = 'scale(1.05) rotate(1.5deg)';
    });
    card.addEventListener('mouseleave', () => {
      card.style.transform = 'scale(1) rotate(0)';
    });
  });

  // Update the search functionality
  searchInput.addEventListener('keydown', debounce(performSearch, 300));

  async function performSearch() {
      const searchValue = searchInput.value.trim();
      suggestionsContainer.innerHTML = '';

      if (searchValue === '') {
          bandCards.forEach(card => card.style.display = 'block');
          return;
      }

      try {
          const response = await fetch(`/search?q=${encodeURIComponent(searchValue)}`);
          if (!response.ok) throw new Error('Search request failed');
          const results = await response.json();
          
          displayResults(results);
      } catch (error) {
          console.error('Error during search:', error);
      }
  }

  function displayResults(results) {
      bandCards.forEach(card => {
          const artistName = card.querySelector('h3').textContent.trim();
          if (results.some(artist => artist.Name === artistName)) {
              card.style.display = 'block';
              addSuggestion(`${artistName} - artist/band`, artistName);
          } else {
              card.style.display = 'none';
          }
      });

      if (results.length === 0) {
          suggestionsContainer.innerHTML = '<div class="suggestion-item">No results found.</div>';
      }
  }

  function addSuggestion(text, value) {
      const suggestionItem = document.createElement('div');
      suggestionItem.textContent = text;
      suggestionItem.classList.add("suggestion-item");
      suggestionItem.addEventListener('click', function() {
          searchInput.value = value;
          suggestionsContainer.innerHTML = '';
          performSearch();
      });
      suggestionsContainer.appendChild(suggestionItem);
  }

  function debounce(func, delay) {
      let timeoutId;
      return function (...args) {
          clearTimeout(timeoutId);
          timeoutId = setTimeout(() => func.apply(this, args), delay);
      };
  }
});
