import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [notes, setNotes] = useState([]);
  const [name, setName] = useState("");
  const [heading, setHeading] = useState("");
  const [content, setContent] = useState("");

  const API = "/api";

  const loadNotes = async () => {
    const res = await axios.get(`${API}/notes`);
    setNotes(res.data);
  };

  useEffect(() => {
    loadNotes();
  }, []);

  const addNote = async () => {
    await axios.post(`${API}/notes`, { name, heading, content });
    setName("");
    setHeading("");
    setContent("");
    loadNotes();
  };

  const deleteNote = async (id) => {
    await axios.delete(`${API}/notes/${id}`);
    loadNotes();
  };

  return (
    <div style={{ padding: 40, fontFamily: "Arial" }}>
      <h2>Notes App</h2>

      <input placeholder="Name" value={name}
        onChange={e => setName(e.target.value)} /><br/><br/>

      <input placeholder="Heading" value={heading}
        onChange={e => setHeading(e.target.value)} /><br/><br/>

      <textarea placeholder="Note" value={content}
        onChange={e => setContent(e.target.value)} /><br/><br/>

      <button onClick={addNote}>Add Note</button>

      <hr/>

      {notes.map(n => (
        <div key={n.id} style={{border:"1px solid #ccc",padding:10,marginTop:10}}>
          <b>{n.heading}</b>
          <p>{n.content}</p>
          <small>{n.name} — {n.created_at}</small><br/>
          <button onClick={()=>deleteNote(n.id)}>Delete</button>
        </div>
      ))}
    </div>
  );
}

export default App;
