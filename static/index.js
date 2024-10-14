document.addEventListener('DOMContentLoaded', (event) => {
    const bandCards = document.querySelectorAll('.band-card');

    // Add hover effect to band cards
    bandCards.forEach(card => {
      card.addEventListener('mouseenter', () => {
        card.style.transform = 'scale(1.05) rotate(1.5deg)';
      });
      card.addEventListener('mouseleave', () => {
        card.style.transform = 'scale(1) rotate(0)';
      });
    });

    // Add click effect to search button
    const searchButton = document.getElementById('search-button');
    searchButton.addEventListener('mousedown', () => {
      searchButton.style.transform = 'translateY(-50%) scale(0.95)';
    });
    searchButton.addEventListener('mouseup', () => {
      searchButton.style.transform = 'translateY(-50%) scale(1)';
    });
    //send post request to search with the bandname as the value to be implemented
    // searchButton.addEventListener('click', ()=>{

    // })
  });

// Search functionality
const searchInput = document.querySelector('#search-input');
const suggestionsContainer = document.querySelector('#suggestions-container');

searchInput.addEventListener('keyup', performSearch)
  function performSearch(){
    const bandCards = document.querySelectorAll('.featured-bands .band-card');
  const searchValue = searchInput.value.toLowerCase().trim();
  suggestionsContainer.innerHTML = ''; // Clear previous suggestions

  if (searchValue === '' && searchValue.length < 2) {
   
    // bandCards.forEach(card => {
    //   card.style.display = 'block';})
      return; // Exit if search input is empty
  }

  bandCards.forEach(card => {
      card.style.display = 'none'; // Hide all cards initially
      let isMatch = false;

      // Search by name
      const bandName = card.querySelector('h3').textContent.trim().toLowerCase();
      if (bandName.includes(searchValue)) {
          addSuggestion(`${bandName} - artist/band`, bandName);
          isMatch = true;
      }

      // Search by members
      const membersStr = card.getAttribute('data-members');
    //   console.log(membersStr);
      if (membersStr) {
          const members = membersStr.split(',').map(m => m.trim());
        //   console.log(members)
          const matchingMembers = members.filter(member => 
              member.toLowerCase().includes(searchValue)
          );

          if (matchingMembers.length > 0) {
            console.log(matchingMembers)
              matchingMembers.forEach(member => {
                member = member.replace(/[\[\]]/g, "");
                  addSuggestion(`${member} - member`, member);
              });
              isMatch = true;
          }
      }

      // Search by creation date
      const creationDate = card.getAttribute('data-creation-date');
      if (creationDate && creationDate.includes(searchValue)) {
          addSuggestion(`Created in ${creationDate} - ${bandName}`, creationDate);
          isMatch = true;
      }

      // Search by first album
      const firstAlbum = card.getAttribute('data-first-album');
      if (firstAlbum && firstAlbum.toLowerCase().includes(searchValue)) {
          addSuggestion(`First album: ${firstAlbum} - ${bandName}`, firstAlbum);
          isMatch = true;
      }

      // Search by locations
      const locationsStr = card.getAttribute('data-locations');
      if (locationsStr) {
          const locations = locationsStr.split(',').map(l => l.trim());
          const matchingLocations = locations.filter(location => 
              location.toLowerCase().includes(searchValue)
          );
          if (matchingLocations.length > 0) {
              matchingLocations.forEach(location => {
                location = location.replace(/[\[\]]/g, "");
                  addSuggestion(`${bandName} - ${location}`, location);
              });
              isMatch = true;
          }
      }

      if (isMatch) {
        card.style.display = 'block';
      }
  });
};

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

//add event listener such that on clearing search input the page is reloaded
searchInput.addEventListener('input', function() {
  if (searchInput.value.trim() === '') {
    window.location.reload(); 
    console.log('input cleared');
  }
});
 