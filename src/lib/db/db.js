import { Level } from "level";

let db;

async function openDB() {
    if (!db) {
        db = new Level("db", { valueEncoding: "json" });
        await db.open();
    }
    return db;
}

async function CreateUsers(ky, data) {
    const db = await openDB();
    try {
        await db.put(ky, data);
        console.log(`Data berhasil disimpan: ${ky}`);
    } catch (error) {
        console.error("Gagal menyimpan data:", error);
    }
}

async function GetUsers(ky) {
    const db = await openDB();
    try {
        return await db.get(ky);
    } catch (error) {
        console.error("Gagal mengambil data:", error);
        return null;
    }
}

async function closeDB() {
    if (db) {
        await db.close();
        db = null;
        console.log("Database ditutup");
    }
}

// Tutup database saat aplikasi ditutup
process.on("SIGINT", async () => {
    await closeDB();
    process.exit(0);
});

// ✅ Ekspor semua fungsi dalam satu objek
export default { CreateUsers, GetUsers, closeDB };
