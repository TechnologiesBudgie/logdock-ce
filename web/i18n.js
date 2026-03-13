/**
 * LogDock i18n Translation Engine
 */

const translations = {
  en: {
    // Nav
    "nav_dashboard": "Dashboard",
    "nav_explorer": "Log Explorer",
    "nav_alerts": "Alerts",
    "nav_settings": "Settings",
    "nav_admin": "Admin",
    "user_role_admin": "Administrator",
    
    // Buttons
    "btn_refresh": "Refresh",
    "btn_export": "Export",
    "btn_live": "Live Tail",
    "btn_signin": "Sign In",
    "btn_save_changes": "Save Changes",
    "btn_cancel": "Cancel",
    "btn_create_rule": "Create Rule",
    "btn_clear_filters": "Clear filters",
    "btn_copy_json": "Copy JSON",
    "btn_done": "Done",

    // Dashboard
    "stat_total": "Total Logs Ingested",
    "stat_errors": "Errors (1m)",
    "stat_rate": "Ingest Rate",
    "stat_disk": "Disk Usage",
    "stat_all_time": "All time",
    "stat_last_60s": "Last 60 seconds",
    "stat_events_min": "Events/min",
    "title_recent_alerts": "Recent Alerts",
    "title_top_sources": "Top Sources",
    "no_data": "No data yet",
    "no_alerts": "No alerts triggered",

    // Explorer
    "lbl_level": "Level:",
    "lbl_range": "Range:",
    "lbl_showing": "Showing",
    "lbl_events": "events",
    "placeholder_search": "Search logs… (e.g. \"error\" or level:error)",
    "col_timestamp": "Timestamp",
    "col_level": "Level",
    "col_source": "Source",
    "col_message": "Message",
    "empty_logs": "No logs found",

    // Login
    "login_welcome": "LogDock CE",
    "lbl_username": "Username",
    "lbl_password": "Password",
    "login_hint": "Default: admin / admin",

    // Alerts
    "alerts_title": "Alert Rules & History",
    "alerts_active": "Active Rules",
    "alerts_triggered": "Triggered Events",
    "no_rules": "No rules configured",
    "no_events": "No triggered events",

    // Settings
    "set_general": "General",
    "set_retention": "Retention",
    "set_security": "Security",
    "set_notifications": "Notifications",
    "set_storage": "Storage",
    "set_instance_name": "Instance Name",
    "set_timezone": "Timezone",
    "lbl_data_dir": "Data Directory",
    "lbl_delete_after": "Retention (days)",
    "hint_delete_after": "Logs older than this are purged",
    "lbl_session_timeout": "Session Duration (min)",
    "lbl_webhook_url": "Slack Webhook URL",
    "msg_settings_saved": "Settings saved",
    "lbl_loading": "Loading…",
    "lbl_partitions": "partitions"
  },
  fr: {
    // Nav
    "nav_dashboard": "Tableau de bord",
    "nav_explorer": "Consultation des logs",
    "nav_alerts": "Alertes",
    "nav_settings": "Paramètres",
    "nav_admin": "Administration",
    "user_role_admin": "Administrateur",

    // Buttons
    "btn_refresh": "Actualiser",
    "btn_export": "Exporter",
    "btn_live": "Flux en direct",
    "btn_signin": "Se connecter",
    "btn_save_changes": "Enregistrer",
    "btn_cancel": "Annuler",
    "btn_create_rule": "Créer la règle",
    "btn_clear_filters": "Effacer les filtres",
    "btn_copy_json": "Copier le JSON",
    "btn_done": "Terminé",

    // Dashboard
    "stat_total": "Total des logs",
    "stat_errors": "Erreurs (1m)",
    "stat_rate": "Débit",
    "stat_disk": "Espace disque",
    "stat_all_time": "Historique",
    "stat_last_60s": "Dernières 60s",
    "stat_events_min": "Événements/min",
    "title_recent_alerts": "Alertes récentes",
    "title_top_sources": "Sources principales",
    "no_data": "Aucune donnée",
    "no_alerts": "Aucune alerte",

    // Explorer
    "lbl_level": "Niveau :",
    "lbl_range": "Période :",
    "lbl_showing": "Affichage de",
    "lbl_events": "événements",
    "placeholder_search": "Rechercher…",
    "col_timestamp": "Horodatage",
    "col_level": "Niveau",
    "col_source": "Source",
    "col_message": "Message",
    "empty_logs": "Aucun log trouvé",

    // Login
    "login_welcome": "LogDock CE",
    "lbl_username": "Identifiant",
    "lbl_password": "Mot de passe",
    "login_hint": "Défaut : admin / admin",

    // Alerts
    "alerts_title": "Alertes",
    "alerts_active": "Règles actives",
    "alerts_triggered": "Événements",
    "no_rules": "Aucune règle",
    "no_events": "Aucun événement",

    // Settings
    "set_general": "Général",
    "set_retention": "Rétention",
    "set_security": "Sécurité",
    "set_notifications": "Notifications",
    "set_storage": "Stockage",
    "set_instance_name": "Nom de l'instance",
    "set_timezone": "Fuseau horaire",
    "lbl_data_dir": "Répertoire de données",
    "lbl_delete_after": "Rétention (jours)",
    "hint_delete_after": "Les logs plus anciens sont supprimés",
    "lbl_session_timeout": "Durée de session (min)",
    "lbl_webhook_url": "URL Webhook Slack",
    "msg_settings_saved": "Paramètres enregistrés",
    "lbl_loading": "Chargement…",
    "lbl_partitions": "partitions"
  }
};

let currentLang = localStorage.getItem('logdock_lang') || 'en';

function initI18n() {
  applyTranslations();
}

function toggleLanguage() {
  currentLang = currentLang === 'en' ? 'fr' : 'en';
  localStorage.setItem('logdock_lang', currentLang);
  applyTranslations();
}

function applyTranslations() {
  document.querySelectorAll('[data-i18n]').forEach(el => {
    const key = el.getAttribute('data-i18n');
    if (translations[currentLang] && translations[currentLang][key]) {
      const val = translations[currentLang][key];
      if ((el.tagName === 'INPUT' || el.tagName === 'TEXTAREA') && el.placeholder) {
        el.placeholder = val;
      } else if (el.tagName === 'TITLE') {
        el.textContent = val;
      } else {
        const span = el.querySelector('span');
        if (span) {
          span.textContent = val;
        } else {
          el.textContent = val;
        }
      }
    }
  });

  const titles = {
    dashboard: translations[currentLang].nav_dashboard,
    explorer: translations[currentLang].nav_explorer,
    alerts: translations[currentLang].nav_alerts,
    settings: translations[currentLang].nav_settings
  };
  
  const pageTitleEl = document.getElementById('page-title');
  if (pageTitleEl && titles[window.currentPage]) {
     pageTitleEl.textContent = titles[window.currentPage];
  }
}

window.initI18n = initI18n;
window.toggleLanguage = toggleLanguage;
window.applyTranslations = applyTranslations;
