document.addEventListener('DOMContentLoaded', function () {
    const searchInput = document.getElementById("searchBox");
    const searchIcon = document.getElementById("searchIcon");
    const clearIcon = document.getElementById("clearIcon");
    const notification = document.getElementById('notification');
    const newsletterLink = document.getElementById('newsletterLink');
    const newsletterModal = document.getElementById('newsletterModal');
    const newsletterForm = document.getElementById('newsletterForm');
    const emailInput = document.getElementById('emailInput');
    const containerMain = document.querySelector('.container-main');

    // Toggle Search Input on Click
    searchIcon.addEventListener("click", function() {
        searchInput.classList.toggle("show");

        if (searchInput.classList.contains("show")) {
            searchInput.focus(); // Focus input when opened
        } else {
            searchInput.blur(); // Remove focus when closed
        }
    });

    // Show Clear Icon When Typing
    searchInput.addEventListener("input", function() {
        if (searchInput.value.length > 0) {
            clearIcon.style.visibility = "visible";
            clearIcon.style.opacity = "1";
        } else {
            clearIcon.style.visibility = "hidden";
            clearIcon.style.opacity = "0";
        }
    });

    // Clear Input When Clicking Cross Icon
    clearIcon.addEventListener("click", function() {
        searchInput.value = "";
        clearIcon.style.visibility = "hidden";
        clearIcon.style.opacity = "0";
        searchInput.classList.toggle("show");
    });

    // RSS link notification
    document.getElementById('rssLink').addEventListener('click', function (event) {
        event.preventDefault(); // Prevent navigation to RSS link

        const rssLink = this.getAttribute("href");

        // Copy to clipboard
        navigator.clipboard.writeText(`${window.location.origin}${rssLink}`)
            .then(() => {
                // Show notification
                notification.textContent = "RSS feed link copied to clipboard!";
                notification.style.opacity = 1;
                notification.style.visibility = 'visible';

                // Hide notification after 3 seconds
                setTimeout(() => {
                    notification.style.opacity = 0;
                    notification.style.visibility = 'hidden';
                    containerMain.classList.remove('blur');
                }, 3000);
            })
            .catch(err => {
                console.error('Failed to copy RSS link: ', err);
            });
    });

    // Show the modal when the "Newsletter" link is clicked
    newsletterLink.addEventListener('click', function (event) {
        event.preventDefault();
        newsletterModal.classList.add('show');
        containerMain.classList.add('blur');
    });

    // Hide the modal when clicking outside of it
    document.addEventListener('click', function (event) {
        if (!newsletterModal.contains(event.target) && !newsletterLink.contains(event.target)) {
            newsletterModal.classList.remove('show');
            containerMain.classList.remove('blur');
        }
    });

    // Handle form submission
    newsletterForm.addEventListener('submit', function (event) {
        event.preventDefault(); // Prevent form from submitting

        const email = emailInput.value.trim();

        if (email) {
            // Replace form content with a success message
            newsletterForm.innerHTML = "<p><h4>Thank you for subscribing!</h4><br> A verification link has been sent to your email.</p>";

            // Close the modal after 2 seconds
            setTimeout(() => {
                newsletterModal.classList.remove('show');
                containerMain.classList.remove('blur');
                emailInput.value = ""; // Clear the input field after submission
                newsletterForm.innerHTML = `
                    <div class="mb-3">
                        <label for="emailInput" class="form-label">Email address</label>
                        <input type="email" class="form-control" id="emailInput" placeholder="Enter your email" required>
                    </div>
                    <button type="submit" class="btn">Subscribe</button>
                `; // Reset form back to original state
            }, 2000);
        } else {
            // Display error message if email is invalid
            alert("Please enter a valid email address.");
            emailInput.focus(); // Focus back on the input field for correction
        }
    });
});