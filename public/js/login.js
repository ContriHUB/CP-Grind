const passwordInput = document.getElementById("password");
const cpasswordInput = document.getElementsByName("cpassword")[0];
const showPasswordButton = document.getElementById("show-password");
const hidePasswordButton = document.getElementById("hide-password");
const showCPasswordButton = document.getElementById("show-cpassword");
const hideCPasswordButton = document.getElementById("hide-cpassword");
const passwordStrengthMessage = document.getElementById("password-strength-message");

showPasswordButton.addEventListener("click", () => {
  passwordInput.type = "text";
  showPasswordButton.style.display = "none";
  hidePasswordButton.style.display = "inline-block";
});

hidePasswordButton.addEventListener("click", () => {
  passwordInput.type = "password";
  showPasswordButton.style.display = "inline-block";
  hidePasswordButton.style.display = "none";
});

showCPasswordButton.addEventListener("click", () => {
  cpasswordInput.type = "text";
  showCPasswordButton.style.display = "none";
  hideCPasswordButton.style.display = "inline-block";
});

hideCPasswordButton.addEventListener("click", () => {
  cpasswordInput.type = "password";
  showCPasswordButton.style.display = "inline-block";
  hideCPasswordButton.style.display = "none";
});

        function checkPasswordStrength() {
          const passwordInput = document.getElementById("password");
          const strengthIndicator = document.getElementById(
            "password-strength-indicator"
          );
          const passwordStrengthMeter = document.getElementById(
            "progress-strength-meter"
          );
          const password = passwordInput.value;

          const result = zxcvbn(password);

          // Set the color indicator and message based on password strength score
       
          switch (result.score) {
            case 0:
              strengthIndicator.innerHTML = "Very Weak";
              passwordStrengthMeter.style.display="block";
              passwordStrengthMeter.style.backgroundColor = "red";
               passwordStrengthMessage.innerHTML = "Very Weak Password";
               passwordStrengthMessage.style.color = "red";
              break;
            case 1:
              strengthIndicator.innerHTML = "Weak";
             passwordStrengthMeter.style.backgroundColor = "orange";
               passwordStrengthMessage.innerHTML = "Weak Password";
               passwordStrengthMessage.style.color = "orange";
              break;
            case 2:
              strengthIndicator.innerHTML = "Fair";
              passwordStrengthMeter.style.backgroundColor = "yellow";
                passwordStrengthMessage.innerHTML = "Fair Password";
                passwordStrengthMessage.style.color = "yellow";
              break;
            case 3:
              strengthIndicator.innerHTML = "Strong";
             passwordStrengthMeter.style.backgroundColor = "green";
               passwordStrengthMessage.innerHTML = "Strong Password";
               passwordStrengthMessage.style.color = "green";
              break;
            case 4:
              strengthIndicator.innerHTML = "Very Strong";
             passwordStrengthMeter.style.backgroundColor = "darkgreen";
               passwordStrengthMessage.innerHTML = "Very Strong Password";
               passwordStrengthMessage.style.color = "darkgreen";
              break;
            default:
              strengthIndicator.innerHTML = "";
          }

          // Update the password strength meter (progress bar)
          const passwordStrengthScore = result.score + 1; // Scores are 0-based, so add 1
          const meterWidth = (passwordStrengthScore / 5) * 100; // Convert to percentage
          passwordStrengthMeter.style.width = `${meterWidth}%`;
          passwordStrengthMeter.setAttribute(
            "aria-valuenow",
            passwordStrengthScore * 20
          ); // Each score is worth 20

           if (result.score >= 0) {
             passwordStrengthMessage.style.display = "block";
           } else {
             passwordStrengthMessage.style.display = "none";
           }
        }

let cnt = 0;
function toggleit(){
    const signup = document.getElementById("sign_up");
    const signin = document.getElementById("sign_in");
    if ((++cnt) & 1){
        signup.style.display = 'none';
        signin.style.display = 'flex';
    }else{
        signin.style.display = 'none';
        signup.style.display = 'flex';
    }
}
