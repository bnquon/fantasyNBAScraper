import { MongoClient } from "mongodb";
import { NextRequest, NextResponse } from 'next/server';

type dataType = {
    email: string;
}

export async function POST(req: NextRequest) {

    const url = process.env.MONGO_URL;
    const data: dataType = await req.json();
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const client = new MongoClient(url!);

    try {
        await client.connect();
        const database = client.db("NBAFantasyProject")
        const collection = database.collection("Emails")
        await collection.insertOne(data)
        return NextResponse.json({ message: "Email saved" })
    } catch (error) {
        return NextResponse.json({ message: "Error saving email" })
    } finally {
        await client.close()
    }
}