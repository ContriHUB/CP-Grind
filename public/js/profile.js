const editButton = document.getElementById("edit-btn");
    const editDropdown = document.getElementById("edit-dropdown");

    editButton.addEventListener("click", function () {
        if (editDropdown.style.display === "none" || editDropdown.style.display === "") {
            editDropdown.style.display = "block";
        } else {
            editDropdown.style.display = "none";
        }
    });

    const dropdownLinks = editDropdown.querySelectorAll("a");

    dropdownLinks.forEach(link => {
        link.addEventListener("click", function() {
            // You can add your own logic here to handle the click action
            // For example, you can navigate to the linked page using JavaScript:
            // window.location.href = link.href;
        });
    });