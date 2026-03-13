import { useEffect, useState, useCallback, useRef } from 'react';

// ─── Types ────────────────────────────────────────────────────────────────────
interface LogRecord {
  id: string; timestamp: string; source: string; level: string;
  message: string; fields?: Record<string, string>;
}
interface Metrics {
  totalIngested: number; errorCount: number; lastMinuteIngest: number;
  diskBytes: number; partitionDays: number;
}
interface AlertRule { name: string; condition: string; target: string; }
interface AlertEvent { rule: string; severity: string; message: string; timestamp: string; }
interface Settings {
  instance_name: string; timezone: string;
  retention: { maxDays: number };
  storage: { dataDir: string };
  security: { sessionHours: number };
  webhooks: { slackUrl: string };
}

type Panel = 'dashboard' | 'explorer' | 'alerts' | 'settings';

// ─── Constants ────────────────────────────────────────────────────────────────
const LEVEL_COLORS: Record<string, { bg: string; text: string; dot: string }> = {
  ERROR:    { bg: '#2d0f0f', text: '#f87171', dot: '#ef4444' },
  WARN:     { bg: '#2d1f0a', text: '#fbbf24', dot: '#f59e0b' },
  INFO:     { bg: '#0a1f2d', text: '#60a5fa', dot: '#3b82f6' },
  DEBUG:    { bg: '#141414', text: '#94a3b8', dot: '#64748b' },
};
const defaultLevel = { bg: '#1a1a2e', text: '#a0aec0', dot: '#718096' };
const lvl = (l: string) => LEVEL_COLORS[(l || '').toUpperCase()] ?? defaultLevel;

const NAV_ITEMS: { id: Panel; label: string; icon: string }[] = [
  { id: 'dashboard', label: 'Dashboard',   icon: '▦' },
  { id: 'explorer',  label: 'Log Explorer', icon: '⌕' },
  { id: 'alerts',    label: 'Alerts',      icon: '⚡' },
  { id: 'settings',  label: 'Settings',    icon: '⚙' },
];

function fmtBytes(b: number) {
  if (b < 1048576) return (b/1024).toFixed(1) + ' KB';
  if (b < 1073741824) return (b/1048576).toFixed(1) + ' MB';
  return (b/1073741824).toFixed(2) + ' GB';
}
function fmtTime(ts: string) {
  try { return new Date(ts).toLocaleTimeString('en-US', { hour12: false, hour:'2-digit', minute:'2-digit', second:'2-digit' }); }
  catch { return ts; }
}

// ─── API Helper ───────────────────────────────────────────────────────────────
async function apiCall<T = any>(path: string, opts?: RequestInit): Promise<T> {
  const token = localStorage.getItem('logdock_token');
  const res = await fetch(path, {
    ...opts,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(opts?.headers ?? {}),
    },
  });
  if (res.status === 401 || res.status === 403) {
    localStorage.removeItem('logdock_token');
    window.location.reload();
    throw new Error('Unauthorized');
  }
  if (res.status === 204) return {} as T;
  return res.json();
}

// ─── Shared UI ────────────────────────────────────────────────────────────────
function Stat({ label, value, sub, color }: { label: string; value: string | number; sub?: string; color?: string }) {
  return (
    <div style={{ background:'#12161f', border:'1px solid #1e2535', borderRadius:8, padding:'20px 24px', flex:1, minWidth:160 }}>
      <div style={{ fontSize:11, color:'#4b5673', textTransform:'uppercase', letterSpacing:'0.1em', marginBottom:4 }}>{label}</div>
      <div style={{ fontSize:28, fontWeight:700, color: color || '#e2e8f0' }}>{value}</div>
      {sub && <div style={{ fontSize:11, color:'#4b5673', marginTop:4 }}>{sub}</div>}
    </div>
  );
}

function Card({ title, children }: { title: string; children: any }) {
  return (
    <div style={{ background:'#0d1220', border:'1px solid #1a2035', borderRadius:8, overflow:'hidden', marginBottom:20 }}>
      <div style={{ padding:'10px 16px', borderBottom:'1px solid #1a2035', fontSize:11, fontWeight:600, color:'#4b5673', textTransform:'uppercase' }}>{title}</div>
      {children}
    </div>
  );
}

// ─── Panel Components ─────────────────────────────────────────────────────────
function DashboardPanel() {
  const [m, setMetrics] = useState<Metrics | null>(null);
  const [logs, setLogs] = useState<LogRecord[]>([]);

  useEffect(() => {
    apiCall('/api/v1/metrics').then(d => setMetrics(d.metrics));
    apiCall('/api/v1/logs/tail').then(d => setLogs(d.rows));
  }, []);

  return (
    <div>
      <div style={{ display:'flex', gap:16, marginBottom:24 }}>
        <Stat label="Total Logs" value={(m?.totalIngested ?? 0).toLocaleString()} color="#60a5fa" />
        <Stat label="Errors" value={(m?.errorCount ?? 0).toLocaleString()} color="#f87171" />
        <Stat label="Ingest Rate" value={(m?.lastMinuteIngest ?? 0)} sub="events/min" color="#4ade80" />
        <Stat label="Disk Usage" value={fmtBytes(m?.diskBytes ?? 0)} sub={`${m?.partitionDays ?? 0} days`} color="#a78bfa" />
      </div>
      <Card title="Recent Events">
        <div style={{ maxHeight:400, overflowY:'auto' }}>
          {logs.map((l, i) => (
            <div key={i} style={{ display:'grid', gridTemplateColumns:'90px 70px 1fr', padding:'8px 16px', borderBottom:'1px solid #0f1420', fontSize:12 }}>
              <span style={{ color:'#475569', fontFamily:'monospace' }}>{fmtTime(l.timestamp)}</span>
              <span style={{ color: lvl(l.level).text, fontWeight:700 }}>{l.level}</span>
              <span style={{ color:'#cbd5e1', overflow:'hidden', textOverflow:'ellipsis', whiteSpace:'nowrap' }}>{l.message}</span>
            </div>
          ))}
        </div>
      </Card>
    </div>
  );
}

function ExplorerPanel() {
  const [q, setQ] = useState('');
  const [rows, setRows] = useState<LogRecord[]>([]);
  const [loading, setLoading] = useState(false);

  const search = useCallback(() => {
    setLoading(true);
    apiCall(`/api/v1/logs/search?q=${encodeURIComponent(q)}&limit=200`)
      .then(d => setRows(d.rows || []))
      .finally(() => setLoading(false));
  }, [q]);

  useEffect(() => { search(); }, []);

  return (
    <div style={{ display:'flex', flexDirection:'column', gap:16 }}>
      <div style={{ background:'#0d1220', border:'1px solid #1a2035', borderRadius:8, padding:12, display:'flex', gap:10 }}>
        <input value={q} onChange={e=>setQ(e.target.value)} onKeyDown={e=>e.key==='Enter'&&search()}
          placeholder="Search logs…" style={{ flex:1, background:'#0a0e18', border:'1px solid #1e2a40', borderRadius:6, color:'#e2e8f0', padding:'8px 12px', outline:'none' }} />
        <button onClick={search} style={{ background:'#1d4ed8', border:'none', color:'white', padding:'8px 20px', borderRadius:6, cursor:'pointer', fontWeight:600 }}>Search</button>
      </div>
      <Card title={`Results (${rows.length})`}>
        <div style={{ height:'calc(100vh - 250px)', overflowY:'auto' }}>
          {loading ? <div style={{ padding:40, textAlign:'center' }}>Loading…</div> : rows.map((l, i) => (
            <div key={i} style={{ display:'grid', gridTemplateColumns:'130px 70px 120px 1fr', padding:'6px 16px', borderBottom:'1px solid #0f1420', fontSize:12 }}>
              <span style={{ color:'#475569', fontFamily:'monospace' }}>{l.timestamp.slice(0,19).replace('T',' ')}</span>
              <span style={{ color: lvl(l.level).text, fontWeight:700 }}>{l.level}</span>
              <span style={{ color:'#64748b' }}>{l.source}</span>
              <span style={{ color:'#cbd5e1' }}>{l.message}</span>
            </div>
          ))}
        </div>
      </Card>
    </div>
  );
}

function AlertsPanel() {
  const [rules, setRules] = useState<AlertRule[]>([]);
  const [history, setHistory] = useState<AlertEvent[]>([]);

  useEffect(() => {
    apiCall('/api/v1/alerts').then(d => setRules(d.rules || []));
    apiCall('/api/v1/alerts/history').then(d => setHistory(d || []));
  }, []);

  return (
    <div style={{ display:'grid', gridTemplateColumns:'1fr 1fr', gap:20 }}>
      <Card title="Active Rules">
        {rules.map((r, i) => (
          <div key={i} style={{ padding:'12px 16px', borderBottom:'1px solid #0f1420' }}>
            <div style={{ fontWeight:600 }}>{r.name}</div>
            <div style={{ fontSize:11, color:'#4b5673', fontFamily:'monospace' }}>{r.condition}</div>
          </div>
        ))}
      </Card>
      <Card title="Alert History">
        {history.reverse().map((e, i) => (
          <div key={i} style={{ padding:'10px 16px', borderBottom:'1px solid #0f1420', display:'flex', justifyContent:'space-between' }}>
            <span style={{ color: e.severity==='high'?'#ef4444':'#f59e0b' }}>{e.rule}</span>
            <span style={{ color:'#4b5673', fontSize:11 }}>{fmtTime(e.timestamp)}</span>
          </div>
        ))}
      </Card>
    </div>
  );
}

function SettingsPanel() {
  const [s, setSettings] = useState<Settings | null>(null);
  const [saved, setSaved] = useState(false);

  useEffect(() => { apiCall('/api/v1/settings').then(setSettings); }, []);

  const save = () => {
    if (!s) return;
    apiCall('/api/v1/settings', { method:'PUT', body: JSON.stringify(s) })
      .then(() => { setSaved(true); setTimeout(()=>setSaved(false), 2000); });
  };

  if (!s) return null;

  return (
    <Card title="System Settings">
      <div style={{ padding:20, display:'flex', flexDirection:'column', gap:16, maxWidth:400 }}>
        <div>
          <label style={{ fontSize:11, color:'#4b5673', textTransform:'uppercase' }}>Instance Name</label>
          <input value={s.instance_name} onChange={e=>setSettings({...s, instance_name:e.target.value})}
            style={{ width:'100%', background:'#0a0e18', border:'1px solid #1e2a40', borderRadius:6, color:'#e2e8f0', padding:8, marginTop:4 }} />
        </div>
        <div>
          <label style={{ fontSize:11, color:'#4b5673', textTransform:'uppercase' }}>Retention (days)</label>
          <input type="number" value={s.retention.maxDays} onChange={e=>setSettings({...s, retention:{maxDays:+e.target.value}})}
            style={{ width:'100%', background:'#0a0e18', border:'1px solid #1e2a40', borderRadius:6, color:'#e2e8f0', padding:8, marginTop:4 }} />
        </div>
        <div>
          <label style={{ fontSize:11, color:'#4b5673', textTransform:'uppercase' }}>Slack Webhook</label>
          <input value={s.webhooks.slackUrl} onChange={e=>setSettings({...s, webhooks:{slackUrl:e.target.value}})}
            style={{ width:'100%', background:'#0a0e18', border:'1px solid #1e2a40', borderRadius:6, color:'#e2e8f0', padding:8, marginTop:4 }} />
        </div>
        <button onClick={save} style={{ background:'#1d4ed8', border:'none', color:'white', padding:'10px', borderRadius:6, cursor:'pointer', fontWeight:600 }}>
          {saved ? '✓ Saved' : 'Save Changes'}
        </button>
      </div>
    </Card>
  );
}

// ─── Main App ─────────────────────────────────────────────────────────────────
export function App() {
  const [token, setToken] = useState<string | null>(localStorage.getItem('logdock_token'));
  const [panel, setPanel] = useState<Panel>('dashboard');

  if (!token) return (
    <div style={{ display:'flex', justifyContent:'center', alignItems:'center', height:'100vh', background:'#070b14' }}>
      <div style={{ background:'#0d1220', border:'1px solid #1a2a40', borderRadius:12, width:320, padding:40, textAlign:'center' }}>
        <h2 style={{ color:'#e2e8f0', marginBottom:20 }}>LogDock CE</h2>
        <input id="u" placeholder="Username" style={{ width:'100%', background:'#070b14', border:'1px solid #1e2a40', borderRadius:6, color:'#e2e8f0', padding:10, marginBottom:10 }} />
        <input id="p" type="password" placeholder="Password" style={{ width:'100%', background:'#070b14', border:'1px solid #1e2a40', borderRadius:6, color:'#e2e8f0', padding:10, marginBottom:20 }} />
        <button onClick={async () => {
          const res = await fetch('/api/v1/auth/login', {
            method:'POST', headers:{'Content-Type':'application/json'},
            body: JSON.stringify({username:(document.getElementById('u') as any).value, password:(document.getElementById('p') as any).value})
          });
          const d = await res.json();
          if (d.token) { localStorage.setItem('logdock_token', d.token); setToken(d.token); }
        }} style={{ width:'100%', background:'#1d4ed8', border:'none', color:'white', padding:12, borderRadius:6, fontWeight:700, cursor:'pointer' }}>Sign In</button>
      </div>
    </div>
  );

  const panels: Record<Panel, JSX.Element> = {
    dashboard: <DashboardPanel />,
    explorer:  <ExplorerPanel />,
    alerts:     <AlertsPanel />,
    settings:   <SettingsPanel />,
  };

  return (
    <div style={{ display:'grid', gridTemplateColumns:'200px 1fr', minHeight:'100vh', background:'#070b14', color:'#e2e8f0', fontFamily:'system-ui' }}>
      <aside style={{ background:'#0a0e18', borderRight:'1px solid #0f1420', padding:10 }}>
        <div style={{ padding:'20px 10px', fontSize:18, fontWeight:800, color:'#3b82f6' }}>LogDock CE</div>
        <nav>
          {NAV_ITEMS.map(item => (
            <button key={item.id} onClick={() => setPanel(item.id)}
              style={{ width:'100%', padding:'12px 15px', marginBottom:4, textAlign:'left', borderRadius:8, border:'none', cursor:'pointer', fontSize:14,
                background: panel === item.id ? '#1e2a40' : 'transparent', color: panel === item.id ? '#e2e8f0' : '#4b5673' }}>
              {item.icon} &nbsp; {item.label}
            </button>
          ))}
        </nav>
        <button onClick={()=>{localStorage.removeItem('logdock_token'); setToken(null);}}
          style={{ position:'absolute', bottom:20, width:160, background:'transparent', border:'1px solid #1e2535', color:'#4b5673', padding:8, borderRadius:6, cursor:'pointer' }}>Sign Out</button>
      </aside>
      <main style={{ padding:30 }}>
        <header style={{ marginBottom:30, display:'flex', justifyContent:'space-between', alignItems:'center' }}>
          <h1 style={{ fontSize:20, fontWeight:700 }}>{NAV_ITEMS.find(n=>n.id===panel)?.label}</h1>
          <div style={{ fontSize:12, color:'#1e2a40' }}>{new Date().toISOString().slice(0,19).replace('T',' ')} UTC</div>
        </header>
        {panels[panel]}
      </main>
    </div>
  );
}
