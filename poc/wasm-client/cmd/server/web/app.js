document.addEventListener("DOMContentLoaded", () => {
    const generateBtn = document.getElementById("generateCsrBtn");
    const outputEl = document.getElementById("output");

    generateBtn.addEventListener("click", async () => {
        try {
            // Assuming `generateCSR` is a function exposed from Go
            // that returns a promise which resolves with the CSR string.
            const csr = await generateCSR();
            outputEl.textContent = csr; // Display the generated CSR
        } catch (err) {
            console.error("Error generating CSR:", err);
            outputEl.textContent = "Failed to generate CSR: " + err.message;
        }
    });
});
