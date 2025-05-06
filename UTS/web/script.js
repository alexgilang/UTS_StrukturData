// Fungsi untuk mengambil data kota dan stasiun dari backend
async function fetchCityData() {
  try {
    const response = await fetch("/api/cities");
    if (!response.ok) {
      throw new Error("Response jaringan bermasalah: " + response.status);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Gagal mengambil data kota:", error);
    return [];
  }
}

// Mendapatkan elemen-elemen HTML
const listContainer = document.getElementById("listContainer");
const searchBox = document.getElementById("searchBox");

// Variabel untuk menyimpan data yang sudah diambil dari backend
let cityData = [];

// Fungsi untuk menampilkan daftar kota dan stasiun dengan filter pencarian
function renderList(filterText = "") {
  listContainer.innerHTML = ""; // Kosongkan isi container dulu
  const filterLower = filterText.toLowerCase(); // Ubah ke huruf kecil untuk pencarian case-insensitive

  // Filter data kota/stasiun berdasarkan filterText
  const filtered = cityData.filter(
    (item) =>
      item.city.toLowerCase().includes(filterLower) || // Cari di nama kota
      item.stations.some((st) => st.toLowerCase().includes(filterLower)) // Cari di list stasiun
  );

  if (filtered.length === 0) {
    // Jika tidak ada hasil, tampilkan pesan
    listContainer.innerHTML =
      '<p style="text-align:center;opacity:0.7;">Tidak ditemukan hasil.</p>';
    return;
  }

  // Jika ada, buat elemen HTML untuk tiap kota dan stasiunnya
  filtered.forEach((item) => {
    const cityDiv = document.createElement("div");
    cityDiv.className = "city-item";
    cityDiv.tabIndex = 0; // Biar bisa fokus untuk aksesibilitas
    cityDiv.setAttribute("role", "listitem");
    cityDiv.innerHTML = `
          <strong>${item.city}</strong>
          <ul class="station-list">
              ${item.stations.map((st) => `<li>${st}</li>`).join("")}
          </ul>
      `;
    listContainer.appendChild(cityDiv);
  });
}

// Event listener pada input pencarian
searchBox.addEventListener("input", (e) => {
  renderList(e.target.value);
});

// Saat halaman dimuat, ambil data dari backend dan render daftar pertama kali
(async () => {
  cityData = await fetchCityData();
  renderList();
})();
