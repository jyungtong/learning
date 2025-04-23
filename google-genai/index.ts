import path from "path";
import {GoogleGenAI, createUserContent,createPartFromUri} from '@google/genai';

(async () => {
const ai = new GoogleGenAI({ apiKey: 'xxx' });

// list all files
// const pager = await ai.files.list({ config: { pageSize: 10 } });
// let page = pager.page;
// const names = [];
// while (true) {
//   for (const f of page) {
//     console.log("  ", f.name);
//     names.push(f.name);
//   }
//   if (!pager.hasNextPage()) break;
//   page = await pager.nextPage();
// }

// delete
// await ai.files.delete({ name: '4u4dli0p30lx' });

// upload
// const myfile = await ai.files.upload({
//   file: path.resolve('./kb.txt'),
//   config: { name: 'gifted-kb' },
// });
// console.log("Uploaded file:", myfile);

// const file = await ai.files.get({ name: 'gifted-kb' });
// console.log("file=", file);

// const result = await ai.models.generateContent({
//   model: "gemini-2.0-flash",
//   contents: createUserContent([
//     createPartFromUri(file.uri, file.mimeType),
//     "\n\n",
//     "what is this about?",
//   ]),
// });
// console.log("result.text=", result.text);
})()
