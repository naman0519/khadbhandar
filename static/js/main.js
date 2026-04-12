


// ================= SAFE DOM LOAD =================
document.addEventListener("DOMContentLoaded", function(){
  console.log("JS LOADED");
});

// ================= WHATSAPP =================
window.toggleWhatsapp = function(){
  let box = document.getElementById("whatsappBox");
  if(box){
    box.style.display = (box.style.display === "block") ? "none" : "block";
  }
};

// ================= HERO SLIDER =================
let slides = document.querySelectorAll(".hero-slide");
let dots = document.querySelectorAll(".dot");
let prev = document.querySelector(".prev");
let next = document.querySelector(".next");

let index = 0;

function showSlide(i){
  if(slides.length === 0) return;

  slides.forEach(s=>s.classList.remove("active"));
  dots.forEach(d=>d.classList.remove("active"));

  if(slides[i]){
    slides[i].classList.add("active");
    if(dots[i]) dots[i].classList.add("active");
  }
}

if(next && prev){
  next.onclick = () => {
    index = (index + 1) % slides.length;
    showSlide(index);
  };

  prev.onclick = () => {
    index = (index - 1 + slides.length) % slides.length;
    showSlide(index);
  };
}

if(slides.length > 0){
  setInterval(()=>{
    index = (index + 1) % slides.length;
    showSlide(index);
  },4000);
}

// ================= PRODUCT SLIDER =================
document.querySelectorAll(".product-slider").forEach(slider => {

  let imgs = slider.querySelectorAll(".product-img");
  let dots = slider.querySelectorAll(".p-dot");

  let i = 0;

  if(imgs.length > 0){
    setInterval(()=>{
      imgs[i].classList.remove("active");
      if(dots[i]) dots[i].classList.remove("active");

      i = (i + 1) % imgs.length;

      imgs[i].classList.add("active");
      if(dots[i]) dots[i].classList.add("active");

    },3000);
  }

});


// ================= FILTER =================
function filterProducts(category, element) {

  let products = document.querySelectorAll(".product-item");
  let buttons = document.querySelectorAll(".filter-btn");

  // active button
  buttons.forEach(btn => btn.classList.remove("active"));
  if (element) element.classList.add("active");

  // filter
  products.forEach(product => {
    let cat = product.getAttribute("data-category");

    if (category === "all" || cat === category) {
      product.style.display = "";
    } else {
      product.style.display = "none";
    }
  });
}


// ================= ESC CLOSE =================
document.addEventListener("keydown", function(e){
  if(e.key === "Escape"){
    let popup = document.getElementById("productPopup");
    if(popup) popup.style.display = "none";
  }
});

// ================= TOGGLE =================
window.toggleService = function(card){
  card.classList.toggle("active");
};

window.toggleTip = function(card){
  card.classList.toggle("active");
};

// ================= MFMB =================

window.startOrder = function(product){
  window.selectedProduct = product.toLowerCase().trim();
  document.getElementById("mfmbPopup").style.display = "flex";
};

window.continueOrder = function(){

  console.log("CLICKED"); // 👈 debug

  if(!window.selectedProduct){
    alert("Product not selected");
    return;
  }

  document.getElementById("mfmbPopup").style.display = "none";

  window.location.href = "/order?product=" + window.selectedProduct;
};



// <!-- FARMER CALCULATOR -->

window.calculateFertilizer = function(){

  let land = parseFloat(document.getElementById("land").value);
  let unit = document.getElementById("unit").value;
  let crop = document.getElementById("crop").value;

  if(!land || land <= 0){
    alert("Enter valid land");
    return;
  }

  // Convert to Bigha
  let bigha = land;
  if(unit === "acre") bigha = land * 2.5;
  if(unit === "killa") bigha = land * 1.25;

  let urea = 0, dap = 0;

  if(crop === "wheat"){ urea = 45; dap = 25; }
  if(crop === "rice"){ urea = 50; dap = 30; }
  if(crop === "mustard"){ urea = 40; dap = 20; }
  if(crop === "maize"){ urea = 55; dap = 30; }
  if(crop === "cotton"){ urea = 60; dap = 35; }
  if(crop === "barley"){ urea = 35; dap = 20; }
  if(crop === "til"){ urea = 30; dap = 15; }

  let totalUrea = urea * bigha;
  let totalDAP = dap * bigha;

  // 💰 PRICE (example)
  let ureaPrice = 6;   // ₹ per kg
  let dapPrice = 25;

  let costUrea = totalUrea * ureaPrice;
  let costDAP = totalDAP * dapPrice;

  let totalCost = costUrea + costDAP;

  document.getElementById("resultBox").innerHTML = `
    <div class="result-card">

      <h5>🌾 Required Fertilizer</h5>

      <p>🌿 Urea: <b>${totalUrea} kg</b></p>
      <p>🌱 DAP: <b>${totalDAP} kg</b></p>

      <hr>

      <h5>💰 Estimated Cost</h5>

      <p>Urea: ₹${costUrea}</p>
      <p>DAP: ₹${costDAP}</p>

      <h4 class="text-success">Total: ₹${totalCost}</h4>

      <a href="/order?product=urea" class="btn btn-success mt-2 w-100">
        Order Fertilizer
      </a>

    </div>
  `;
}

// UPDATE SCRIPT

function updateStock(id, product){

  console.log("CLICK:", id, product); // debug

  let stock = document.getElementById("stock_" + id).value;

  fetch("/admin/update-stock", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded"
    },
    body: "product=" + product + "&stock=" + stock
  })
  .then(res => res.json())
  .then(data => {
    alert(product + " updated ✅");
    location.reload();
  });

}

// SHOW OR HIDE PASSWORD

function togglePassword() {
  var x = document.querySelector('input[name="password"]');
  if (x.type === "password") {
    x.type = "text";
  } else {
    x.type = "password";
  }
}


// LOADING BUTTON 
function showLoading(btn) {
  btn.innerText = "Logging in...";
}



window.addEventListener("scroll", function () {
  const navbar = document.querySelector(".navbar");

  if (window.scrollY > 50) {
    navbar.classList.add("scrolled");
  } else {
    navbar.classList.remove("scrolled");
  }
});




const faders = document.querySelectorAll(".fade-in");

window.addEventListener("scroll", () => {
  faders.forEach(el => {
    const rect = el.getBoundingClientRect();

    if (rect.top < window.innerHeight - 100) {
      el.classList.add("show");
    }
  });
});



// ---------VIEW MORE--------------//

function toggleProducts() {
  const extra = document.getElementById("extraProducts");
  const btn = document.getElementById("toggleBtn");

  if (extra.style.display === "none" || extra.style.display === "") {
    extra.style.display = "block";
    btn.innerText = "View Less";
  } else {
    extra.style.display = "none";
    btn.innerText = "View More";
  }
}


// ================= ORDER VALIDATION =================

let form = document.getElementById("orderForm");

if(form){
  form.addEventListener("submit", function(e){

    let name = document.getElementById("name").value.trim();
    let phone = document.getElementById("phone").value.trim();
    let quantity = document.getElementById("quantity").value;

    //  empty check
    if(name === "" || phone === "" || quantity === ""){
      alert("⚠ Fill all fields");
      e.preventDefault();
      return;
    }

    //  phone check
    if(!/^[0-9]{10}$/.test(phone)){
      alert("📱 Enter valid 10 digit number");
      e.preventDefault();
      return;
    }

    // quantity check
    if(quantity <= 0){
      alert(" Quantity must be greater than 0");
      e.preventDefault();
      return;
    }

    // MAX LIMIT
    if(quantity > 30){
      alert(" Maximum 30 bags allowed");
      e.preventDefault();
      return;
    }

  });
}


// ================= LIVE STOCK CHECK =================

let qtyInput = document.getElementById("quantity");
let stockValue = document.getElementById("stockValue");
let stockMsg = document.getElementById("stockMsg");

if(qtyInput && stockValue){

  let stock = parseInt(stockValue.value);

  qtyInput.addEventListener("input", function(){

    let qty = parseInt(qtyInput.value);

    if(!qty){
      stockMsg.innerText = "";
      return;
    }

    if(stock === 0){
      stockMsg.innerText = " Out of stock";
      stockMsg.style.color = "red";
      return;
    }

    if(qty > stock){
      stockMsg.innerText = "⚠ Only " + stock + " available";
      stockMsg.style.color = "orange";
    } else {
      stockMsg.innerText = " Available";
      stockMsg.style.color = "green";
    }

  });
}




// ==========ORDER FLOW WITH MFMB==========



function startOrder(product){
  selectedProduct = product;
  document.getElementById("mfmbPopup").style.display = "flex";
}

function continueOrder(){
  if(!selectedProduct){
    alert("Product missing");
    return;
  }
  window.location.href = "/order?product=" + selectedProduct;
}



// ================= FILTER + SEARCH COMBINED =================



  const buttons = document.querySelectorAll(".filter-btn");
  const products = document.querySelectorAll(".product-item");
  const searchInput = document.getElementById("searchInput");

  let currentFilter = "all";

  // FILTER CLICK
  buttons.forEach(btn => {
    btn.addEventListener("click", function () {

      currentFilter = this.getAttribute("data-filter");

      buttons.forEach(b => b.classList.remove("active"));
      this.classList.add("active");

      applyFilterAndSearch();
    });
  });

  // SEARCH
  if (searchInput) {
    searchInput.addEventListener("keyup", function () {
      applyFilterAndSearch();
    });
  }

  function applyFilterAndSearch() {
    const searchValue = searchInput ? searchInput.value.toLowerCase() : "";

    let firstMatch = null;

    products.forEach(item => {
      const category = item.getAttribute("data-category");
      const name = item.innerText.toLowerCase();

      const matchCategory = (currentFilter === "all" || category === currentFilter);
      const matchSearch = name.includes(searchValue);

      if (matchCategory && matchSearch) {
        item.style.display = "block";

        if (!firstMatch && searchValue !== "") {
          firstMatch = item;
        }

      } else {
        item.style.display = "none";
      }
    });

    // scroll
    if (firstMatch) {
      firstMatch.scrollIntoView({ behavior: "smooth", block: "center" });
    }
  }








// ================= POPUP =================

window.selectedProduct = "";

// OPEN POPUP
window.openPopup = function(product) {

  alert("BUTTON CLICKED: " + product); // debug

  const popup = document.getElementById("productPopup");

  window.selectedProduct = product;

  popup.style.display = "flex";

  document.getElementById("popupTitle").innerText = product.toUpperCase();
  document.getElementById("popupPrice").innerText = "₹266";
  document.getElementById("popupDesc").innerText = "Test fertilizer";
  document.getElementById("popupImage").src = "/static/images/urea1.jpg";
}

window.closePopup = function() {
  document.getElementById("productPopup").style.display = "none";
}


// SHOW PRODUCTS BASED ON CROP SELECTION

function showCropProducts() {

  const selectedCrop = document.getElementById("cropSelect").value;

  const products = document.querySelectorAll(".product-item");

  products.forEach(product => {

    const crops = product.getAttribute("data-crop");

    if (!selectedCrop) {
      product.style.display = "block";
      return;
    }

    if (crops && crops.includes(selectedCrop)) {
      product.style.display = "block";
    } else {
      product.style.display = "none";
    }

  });

}
