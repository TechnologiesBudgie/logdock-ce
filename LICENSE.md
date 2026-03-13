================================================================================
TECHNOLOGIES BUDGIE
LOGDOCK COMMUNITY EDITION LICENSE AGREEMENT (LDCEL)
Version 1.0 — January 1, 2025

TECHNOLOGIES BUDGIE
CONTRAT DE LICENCE — LOGDOCK ÉDITION COMMUNAUTAIRE (LDLEC)
Version 1.0 — 1er janvier 2025
================================================================================

Licensing & Contact / Licence et contact :
Technologies Budgie
Email  : budgie@mailfence.com
Website: https://technologiesbudgie.page.dev

To inquire about upgrading to the LogDock Enterprise Edition, or to report a
license violation, contact: budgie@mailfence.com

Pour obtenir une licence Entreprise ou signaler une violation : budgie@mailfence.com

================================================================================
SECTION 0 — PREAMBLE / PRÉAMBULE
================================================================================

--- ENGLISH ---

Technologies Budgie ("the Company," "Licensor," "we," or "us") is a Québec-based
software company. LogDock is a log and data storage coordination platform
developed and maintained by Technologies Budgie. It includes storage management,
automated triage, AI-assisted pipelines, compression engines, analytical
database layers, and related tooling.

This LogDock Community Edition License Agreement ("Agreement" or "LDCEL")
governs only the Community Edition of LogDock. It is designed to permit broad
community use — individual developers, researchers, students, non-profits, and
organizations operating LogDock for their own internal purposes — while
protecting the Company's ability to sustain development and maintain a viable
commercial offering.

This Agreement prevents third parties from offering competing hosted or managed
services built substantially upon LogDock without a commercial license, and
protects the trademarks "Technologies Budgie" and "LogDock."

Both the English and French texts of this Agreement are legally valid. In any
proceeding before a Québec court, the French text prevails to the extent
required by the Charter of the French Language (RLRQ c C-11). In all other
jurisdictions, the English text prevails.

BY DOWNLOADING, INSTALLING, COPYING, RUNNING, MODIFYING, OR OTHERWISE USING
LOGDOCK COMMUNITY EDITION, YOU AGREE TO ALL TERMS OF THIS AGREEMENT. IF YOU DO
NOT AGREE, YOU MUST NOT USE THE SOFTWARE.

--- FRANÇAIS ---

Technologies Budgie (« la Société », « le Concédant », « nous ») est une société
de logiciels basée au Québec. LogDock est une plateforme de coordination de
stockage de journaux et de données développée et maintenue par Technologies
Budgie. Elle comprend la gestion du stockage, le triage automatisé, des
pipelines assistés par intelligence artificielle, des moteurs de compression,
des couches de bases de données analytiques et des outils connexes.

Le présent Contrat de Licence LogDock — Édition Communautaire (« Contrat » ou
« LDLEC ») régit uniquement l'Édition Communautaire de LogDock. Il est conçu
pour permettre une utilisation communautaire large — développeurs individuels,
chercheurs, étudiants, organisations à but non lucratif et organisations
exploitant LogDock à des fins internes — tout en protégeant la capacité de la
Société à soutenir le développement et à maintenir une offre commerciale viable.

Le présent Contrat interdit aux tiers d'offrir des services hébergés ou gérés
concurrents construits substantiellement sur LogDock sans licence commerciale,
et protège les marques de commerce « Technologies Budgie » et « LogDock ».

Les textes anglais et français du présent Contrat ont tous deux valeur juridique.
Dans toute procédure devant un tribunal québécois, le texte français prévaut dans
la mesure requise par la Charte de la langue française (RLRQ c C-11). Dans toutes
les autres juridictions, le texte anglais prévaut.

EN TÉLÉCHARGEANT, INSTALLANT, COPIANT, EXÉCUTANT, MODIFIANT OU UTILISANT DE
QUELQUE AUTRE MANIÈRE LOGDOCK ÉDITION COMMUNAUTAIRE, VOUS ACCEPTEZ L'ENSEMBLE
DES CONDITIONS DU PRÉSENT CONTRAT. SI VOUS N'ACCEPTEZ PAS, VOUS NE DEVEZ PAS
UTILISER LE LOGICIEL.

================================================================================
SECTION 1 — DEFINITIONS / DÉFINITIONS
================================================================================

--- ENGLISH ---

1.1  "Software" means the LogDock Community Edition source code, compiled
     binaries, Docker images, container definitions, configuration files,
     scripts, documentation, and all associated files distributed by
     Technologies Budgie under this Agreement.

1.2  "Licensor" means Technologies Budgie.

1.3  "You" or "Licensee" means any individual or legal entity exercising
     rights under this Agreement.

1.4  "Affiliate" means any entity that controls, is controlled by, or is under
     common control with a party, where "control" means ownership of more than
     fifty percent (50%) of the voting interests.

1.5  "Contribution" means any original work of authorship submitted to the
     Licensor for inclusion in the Software, including source code, bug fixes,
     documentation, tests, and configuration.

1.6  "Contributor" means any individual or legal entity that submits a
     Contribution.

1.7  "Derivative Work" means any work that is based upon the Software, from
     which the Software or a substantial portion thereof has been used,
     transformed, translated, adapted, or built upon, in any form.

1.8  "Distribute" or "Distribution" means to make the Software or a Derivative
     Work available to any third party by any means, including but not limited
     to publishing source code, offering compiled binaries, making available
     Docker images, publishing to package registries, or providing access via
     network download.

1.9  "Enterprise Edition" means the separately licensed version of LogDock
     made available by Technologies Budgie under a commercial license agreement
     (the LogDock Enterprise Edition License Agreement, "LDEELA"), which
     includes additional features, capabilities, support, and rights not
     available under this Agreement.

1.10 "Enterprise Features" means any functionality, module, algorithm, or
     component designated by Technologies Budgie as exclusive to the Enterprise
     Edition, including but not limited to: advanced AI Triage Integration
     beyond the CE-tier threshold, DuckDB query optimizations, multi-tenant
     SaaS management interfaces, advanced compression algorithm implementations,
     enterprise SSO/LDAP integrations, high-availability clustering modules,
     and any feature marked in the source code or documentation with the tag
     [EE-ONLY] or equivalent.

1.11 "AI Triage Integration" means any feature, module, subsystem, algorithm,
     or interface within the Software that uses machine learning, large language
     models, neural networks, heuristic scoring engines, or other artificial
     intelligence or automated classification techniques to classify, score,
     route, deduplicate, prioritize, or otherwise process log entries, storage
     events, backup jobs, alerts, or associated metadata.

1.12 "SaaS Use" means operating the Software, or any Derivative Work based
     thereon, as a hosted, managed, or cloud-delivered service accessible by
     third parties over a network, including but not limited to: multi-tenant
     log management services, hosted backup coordination platforms, managed
     storage analytics services, and any service for which the primary value
     delivered to end users derives substantially from LogDock functionality.

1.13 "Internal Use" means deployment of the Software solely for the internal
     business operations of the Licensee and its Affiliates, where no access to
     the Software's functionality is provided to any third party, whether for
     compensation or otherwise.

1.14 "Minor Commercial Use" means Internal Use within an organization that
     derives revenue from activities other than operating the Software as a
     service, provided the total number of users (employees and contractors)
     accessing the Software does not exceed one hundred (100), and provided the
     use does not constitute SaaS Use as defined herein.

1.15 "Storage Backend" means any storage subsystem with which LogDock
     integrates or manages, including but not limited to: Linux LVM (Logical
     Volume Manager) configurations, NAS (Network-Attached Storage) systems,
     ZFS (Zettabyte File System) pools and datasets, block storage devices, and
     object storage endpoints.

1.16 "Container Deployment" means any deployment of the Software using
     containerization technology including Docker, Podman, containerd,
     Kubernetes, OpenShift, Nomad, or any equivalent container orchestration or
     runtime environment.

1.17 "Go Binary" means any compiled executable produced from the Software's Go
     language source code, distributed as a standalone binary or as part of a
     Container Deployment.

1.18 "Source Code" means the human-readable form of the Software, including all
     modules, packages, configuration files, build scripts, and associated
     documentation as distributed by the Licensor.

1.19 "Object Code" means any compiled, assembled, minified, or otherwise
     machine-readable form of the Software, including Go binaries and any
     bundled or embedded runtime artifacts.

1.20 "Trademark" means the names "Technologies Budgie" and "LogDock," and any
     associated logos, slogans, or trade dress owned by Technologies Budgie,
     whether or not registered.

1.21 "License Notice" means the copyright and license notice text specified in
     Section 4 of this Agreement, which must be included in all copies and
     Distributions of the Software.

1.22 "Competing Service" means any service, product, or platform offered to
     third parties that provides substantially similar functionality to LogDock,
     including log ingestion, storage management, automated triage, or data
     pipeline coordination, where such service is built upon, derived from, or
     incorporates a substantial portion of the Software.

--- FRANÇAIS ---

1.1  « Logiciel » désigne le code source de LogDock Édition Communautaire, les
     binaires compilés, les images Docker, les définitions de conteneurs, les
     fichiers de configuration, les scripts, la documentation et tous les
     fichiers associés distribués par Technologies Budgie dans le cadre du
     présent Contrat.

1.2  « Concédant » désigne Technologies Budgie.

1.3  « Vous » ou « Licencié » désigne toute personne physique ou morale
     exerçant des droits en vertu du présent Contrat.

1.4  « Affilié » désigne toute entité qui contrôle, est contrôlée par, ou est
     sous contrôle commun avec une partie, où « contrôle » signifie la
     propriété de plus de cinquante pour cent (50 %) des droits de vote.

1.5  « Contribution » désigne toute œuvre originale soumise au Concédant pour
     inclusion dans le Logiciel, y compris le code source, les corrections de
     bogues, la documentation, les tests et la configuration.

1.6  « Contributeur » désigne toute personne physique ou morale qui soumet une
     Contribution.

1.7  « Œuvre Dérivée » désigne toute œuvre basée sur le Logiciel, à partir
     duquel le Logiciel ou une partie substantielle de celui-ci a été utilisé,
     transformé, traduit, adapté ou sur lequel on s'est appuyé, sous quelque
     forme que ce soit.

1.8  « Distribuer » ou « Distribution » désigne le fait de mettre le Logiciel
     ou une Œuvre Dérivée à la disposition de tout tiers par quelque moyen que
     ce soit, incluant notamment la publication du code source, l'offre de
     binaires compilés, la mise à disposition d'images Docker, la publication
     dans des registres de paquets, ou la fourniture d'accès par téléchargement
     réseau.

1.9  « Édition Entreprise » désigne la version séparément licenciée de LogDock
     mise à disposition par Technologies Budgie dans le cadre d'un contrat de
     licence commerciale (le Contrat de Licence LogDock Édition Entreprise,
     « LDLEE »), qui inclut des fonctionnalités, capacités, support et droits
     supplémentaires non disponibles dans le cadre du présent Contrat.

1.10 « Fonctionnalités Entreprise » désigne toute fonctionnalité, module,
     algorithme ou composant désigné par Technologies Budgie comme exclusif à
     l'Édition Entreprise, incluant notamment : l'Intégration IA Triage avancée
     au-delà du seuil de l'EC, les optimisations de requêtes DuckDB, les
     interfaces de gestion SaaS multi-locataires, les implémentations avancées
     d'algorithmes de compression, les intégrations SSO/LDAP entreprise, les
     modules de clustering haute disponibilité, et toute fonctionnalité marquée
     dans le code source ou la documentation avec l'étiquette [EE-ONLY] ou
     équivalent.

1.11 « Intégration IA Triage » désigne toute fonctionnalité, module, sous-
     système, algorithme ou interface au sein du Logiciel qui utilise
     l'apprentissage automatique, les grands modèles de langage, les réseaux de
     neurones, les moteurs de notation heuristique, ou d'autres techniques
     d'intelligence artificielle ou de classification automatisée pour
     classifier, noter, acheminer, dédupliquer, prioriser ou traiter de toute
     autre manière les entrées de journaux, les événements de stockage, les
     tâches de sauvegarde, les alertes ou les métadonnées associées.

1.12 « Usage SaaS » désigne l'exploitation du Logiciel, ou de toute Œuvre
     Dérivée, en tant que service hébergé, géré ou livré en nuage, accessible
     par des tiers via un réseau, incluant notamment : les services de gestion
     de journaux multi-locataires, les plateformes de coordination de sauvegarde
     hébergées, les services d'analyse de stockage gérés, et tout service pour
     lequel la valeur principale fournie aux utilisateurs finaux provient
     substantiellement des fonctionnalités de LogDock.

1.13 « Usage Interne » désigne le déploiement du Logiciel uniquement pour les
     opérations commerciales internes du Licencié et de ses Affiliés, sans
     qu'aucun accès aux fonctionnalités du Logiciel ne soit fourni à un tiers,
     que ce soit à titre onéreux ou autrement.

1.14 « Usage Commercial Mineur » désigne l'Usage Interne au sein d'une
     organisation qui tire des revenus d'activités autres que l'exploitation du
     Logiciel en tant que service, à condition que le nombre total d'utilisateurs
     (employés et contractuels) accédant au Logiciel ne dépasse pas cent (100),
     et à condition que l'usage ne constitue pas un Usage SaaS tel que défini
     aux présentes.

1.15 « Backend de Stockage » désigne tout sous-système de stockage avec lequel
     LogDock s'intègre ou qu'il gère, incluant notamment : les configurations
     Linux LVM (Gestionnaire de Volumes Logiques), les systèmes NAS (Stockage en
     Réseau), les pools et jeux de données ZFS (Système de Fichiers Zettaoctet),
     les dispositifs de stockage en blocs et les points de terminaison de
     stockage d'objets.

1.16 « Déploiement en Conteneur » désigne tout déploiement du Logiciel utilisant
     une technologie de conteneurisation incluant Docker, Podman, containerd,
     Kubernetes, OpenShift, Nomad, ou tout environnement équivalent d'exécution
     ou d'orchestration de conteneurs.

1.17 « Binaire Go » désigne tout exécutable compilé produit à partir du code
     source Go du Logiciel, distribué en tant que binaire autonome ou dans le
     cadre d'un Déploiement en Conteneur.

1.18 « Code Source » désigne la forme lisible par l'homme du Logiciel, incluant
     tous les modules, paquets, fichiers de configuration, scripts de
     compilation et documentation associée tels que distribués par le Concédant.

1.19 « Code Objet » désigne toute forme compilée, assemblée, minifiée ou
     autrement lisible par une machine du Logiciel, incluant les Binaires Go et
     tout artefact d'exécution groupé ou intégré.

1.20 « Marque de Commerce » désigne les noms « Technologies Budgie » et
     « LogDock », ainsi que tout logo, slogan ou habillage commercial associé
     appartenant à Technologies Budgie, qu'ils soient enregistrés ou non.

1.21 « Avis de Licence » désigne le texte d'avis de droit d'auteur et de
     licence spécifié à la Section 4 du présent Contrat, qui doit être inclus
     dans toutes les copies et Distributions du Logiciel.

1.22 « Service Concurrent » désigne tout service, produit ou plateforme offert
     à des tiers qui fournit des fonctionnalités substantiellement similaires à
     LogDock, incluant l'ingestion de journaux, la gestion du stockage, le
     triage automatisé, ou la coordination de pipelines de données, lorsque ce
     service est construit sur, dérivé de, ou incorpore une partie substantielle
     du Logiciel.

================================================================================
SECTION 2 — GRANT OF RIGHTS / DROITS ACCORDÉS
================================================================================

--- ENGLISH ---

Subject to full compliance with all terms and conditions of this Agreement,
Technologies Budgie hereby grants You a worldwide, non-exclusive, royalty-free,
non-transferable, non-sublicensable license to:

2.1  PERSONAL AND EDUCATIONAL USE
     Use, run, copy, and study the Software for personal, educational,
     research, and non-commercial purposes without restriction, subject only to
     the obligations set forth in Sections 3 and 4.

2.2  INTERNAL USE AND MINOR COMMERCIAL USE
     Deploy and operate the Software for Internal Use or Minor Commercial Use,
     including running LogDock as part of your organization's internal
     infrastructure, internal data pipelines, internal log management, and
     internal storage coordination. This right explicitly includes:

     (a) Running LogDock on physical servers, virtual machines, or in a private
         cloud environment controlled exclusively by You or your Affiliates;
     (b) Integrating LogDock with Storage Backends including LVM, NAS, and ZFS
         systems operated for Your own data;
     (c) Using LogDock's Container Deployment capabilities, including Docker and
         Go Binary deployments, for internal infrastructure;
     (d) Using the CE-tier AI Triage Integration features as provided in the
         Community Edition for internal log processing and triage;
     (e) Connecting LogDock to internal databases, including DuckDB instances,
         for internal analytical workloads.

2.3  MODIFICATION
     Modify the Software and create Derivative Works, provided that:

     (a) All modified files carry prominent notices stating that You changed the
         files and the date of such changes;
     (b) Any Derivative Work you Distribute is licensed under this Agreement in
         its entirety, with no additional restrictions imposed beyond those
         already present herein;
     (c) You do not incorporate, enable, or expose Enterprise Features in any
         modified version distributed under this Agreement.

2.4  DISTRIBUTION OF SOURCE CODE
     Distribute copies of the Software's Source Code, in original or modified
     form, provided that:

     (a) You include a complete, unmodified copy of this Agreement with every
         Distribution;
     (b) You include the License Notice required by Section 4;
     (c) You make the complete corresponding Source Code available under the
         terms of this Agreement at no additional charge;
     (d) You do not remove, alter, or obscure any copyright, patent, trademark,
         or attribution notices contained in the Software;
     (e) You comply with all conditions of Section 5 regarding redistribution.

2.5  DISTRIBUTION OF OBJECT CODE AND BINARIES
     Distribute Object Code, compiled Go Binaries, or Container images of the
     Software, provided that:

     (a) You make the complete corresponding Source Code available under this
         Agreement, either by including it with the Distribution or by providing
         a written offer, valid for at least three (3) years, to supply the
         Source Code upon request at no charge beyond reasonable reproduction
         costs;
     (b) You include this Agreement in its entirety with the Distribution;
     (c) You include the License Notice required by Section 4;
     (d) The Binaries or containers do not include or expose Enterprise Features.

2.6  NON-PROFIT AND COMMUNITY HOSTING
     Operate the Software to provide services to a defined, non-commercial
     community (such as an open-source project's developer community, a
     university department, or a non-profit organization's members), provided
     that:

     (a) No revenue or compensation is received in connection with the operation
         of such service;
     (b) The service is not positioned as, marketed as, or functionally
         equivalent to a commercial LogDock-powered service;
     (c) Access is restricted to the defined community and not offered to the
         general public on a commercial basis.

--- FRANÇAIS ---

Sous réserve du respect intégral de l'ensemble des termes et conditions du
présent Contrat, Technologies Budgie vous accorde par les présentes une licence
mondiale, non exclusive, libre de redevances, non transférable et non
sous-licenciable pour :

2.1  USAGE PERSONNEL ET ÉDUCATIF
     Utiliser, exécuter, copier et étudier le Logiciel à des fins personnelles,
     éducatives, de recherche et non commerciales sans restriction, sous réserve
     uniquement des obligations énoncées aux Sections 3 et 4.

2.2  USAGE INTERNE ET USAGE COMMERCIAL MINEUR
     Déployer et exploiter le Logiciel pour un Usage Interne ou un Usage
     Commercial Mineur, incluant l'exécution de LogDock dans le cadre de
     l'infrastructure interne de votre organisation, des pipelines de données
     internes, de la gestion interne des journaux et de la coordination du
     stockage interne. Ce droit inclut explicitement :

     (a) L'exécution de LogDock sur des serveurs physiques, des machines
         virtuelles ou dans un environnement de nuage privé contrôlé
         exclusivement par Vous ou vos Affiliés ;
     (b) L'intégration de LogDock avec des Backends de Stockage incluant les
         systèmes LVM, NAS et ZFS exploités pour vos propres données ;
     (c) L'utilisation des capacités de Déploiement en Conteneur de LogDock,
         incluant les déploiements Docker et Binaire Go, pour l'infrastructure
         interne ;
     (d) L'utilisation des fonctionnalités d'Intégration IA Triage de niveau EC
         telles que fournies dans l'Édition Communautaire pour le traitement et
         le triage internes des journaux ;
     (e) La connexion de LogDock à des bases de données internes, incluant des
         instances DuckDB, pour des charges de travail analytiques internes.

2.3  MODIFICATION
     Modifier le Logiciel et créer des Œuvres Dérivées, à condition que :

     (a) Tous les fichiers modifiés portent des avis bien visibles indiquant que
         Vous avez modifié les fichiers et la date de ces modifications ;
     (b) Toute Œuvre Dérivée que vous Distribuez soit licenciée sous le présent
         Contrat dans son intégralité, sans restrictions supplémentaires au-delà
         de celles déjà présentes dans les présentes ;
     (c) Vous n'intégriez pas, n'activiez pas et n'exposiez pas les
         Fonctionnalités Entreprise dans une version modifiée distribuée dans le
         cadre du présent Contrat.

2.4  DISTRIBUTION DU CODE SOURCE
     Distribuer des copies du Code Source du Logiciel, sous forme originale ou
     modifiée, à condition que :

     (a) Vous incluiez une copie complète et non modifiée du présent Contrat
         avec chaque Distribution ;
     (b) Vous incluiez l'Avis de Licence requis par la Section 4 ;
     (c) Vous mettiez à disposition le Code Source correspondant complet dans le
         cadre du présent Contrat sans frais supplémentaires ;
     (d) Vous ne supprimiez pas, ne modifiiez pas et n'obscurcissiez pas les
         avis de droit d'auteur, de brevet, de marque de commerce ou
         d'attribution contenus dans le Logiciel ;
     (e) Vous respectiez toutes les conditions de la Section 5 concernant la
         redistribution.

2.5  DISTRIBUTION DU CODE OBJET ET DES BINAIRES
     Distribuer le Code Objet, les Binaires Go compilés, ou les images de
     conteneurs du Logiciel, à condition que :

     (a) Vous mettiez à disposition le Code Source correspondant complet dans le
         cadre du présent Contrat, soit en l'incluant avec la Distribution, soit
         en fournissant une offre écrite, valable au moins trois (3) ans, de
         fournir le Code Source sur demande sans frais au-delà des coûts
         raisonnables de reproduction ;
     (b) Vous incluiez le présent Contrat dans son intégralité avec la
         Distribution ;
     (c) Vous incluiez l'Avis de Licence requis par la Section 4 ;
     (d) Les Binaires ou conteneurs n'incluent pas et n'exposent pas les
         Fonctionnalités Entreprise.

2.6  HÉBERGEMENT NON LUCRATIF ET COMMUNAUTAIRE
     Exploiter le Logiciel pour fournir des services à une communauté définie et
     non commerciale (telle que la communauté de développeurs d'un projet
     open-source, un département universitaire ou les membres d'une organisation
     à but non lucratif), à condition que :

     (a) Aucun revenu ou compensation ne soit perçu en lien avec l'exploitation
         d'un tel service ;
     (b) Le service ne soit pas positionné comme, commercialisé comme, ou
         fonctionnellement équivalent à un service commercial propulsé par
         LogDock ;
     (c) L'accès soit limité à la communauté définie et ne soit pas offert au
         grand public sur une base commerciale.

================================================================================
SECTION 3 — RESTRICTIONS / RESTRICTIONS
================================================================================

--- ENGLISH ---

The rights granted in Section 2 are subject to the following restrictions.
Violation of any restriction in this Section constitutes a material breach of
this Agreement and automatically terminates Your rights under Section 11.

3.1  NO SAAS USE WITHOUT ENTERPRISE LICENSE
     You may not use the Software, or any Derivative Work, to provide a
     Competing Service or to operate any SaaS Use accessible to third parties,
     without first obtaining a LogDock Enterprise Edition license from
     Technologies Budgie. This restriction applies regardless of:

     (a) Whether You charge for the service or offer it free of charge;
     (b) Whether the service is marketed or branded under a name other than
         "LogDock";
     (c) Whether You have modified the Software prior to deployment;
     (d) Whether the Software is deployed as a Docker container, Go binary,
         cloud instance, or any other form;
     (e) Whether the service is described as a "wrapper," "layer," "connector,"
         "integration," or any other functional description that obscures its
         reliance on LogDock.

     For the avoidance of doubt: operating LogDock as a shared internal platform
     accessible to employees and contractors of a single organization and its
     direct Affiliates does NOT constitute SaaS Use, provided no third party
     (including customers, partners, or the general public) has access to the
     LogDock-powered functionality.

3.2  NO ENTERPRISE FEATURE EXTRACTION OR DISTRIBUTION
     You may not, under this Agreement:

     (a) Extract, copy, port, reimplement, or distribute any Enterprise Feature
         or any code implementing an Enterprise Feature;
     (b) Modify the Software to remove feature-gating mechanisms that restrict
         access to Enterprise Features, whether such mechanisms are implemented
         via license key checks, build tags, compile-time flags, runtime checks,
         or any other technical means;
     (c) Distribute any version of the Software in which Enterprise Features
         have been unlocked, enabled, or made accessible without a valid
         Enterprise Edition license;
     (d) Reverse-engineer Enterprise Features for the purpose of reimplementing
         them in a competing product, whether or not You have access to the
         corresponding source code.

3.3  TRADEMARK RESTRICTIONS
     You may not:

     (a) Use the Trademarks "Technologies Budgie" or "LogDock" in the name,
         branding, marketing, or documentation of any product, service, or
         organization that is not Technologies Budgie, without prior written
         consent from Technologies Budgie;
     (b) Represent that Your Derivative Work or service is the official LogDock
         product or is endorsed, sponsored, or approved by Technologies Budgie;
     (c) Register or attempt to register any trademark, domain name, social
         media handle, or business name that incorporates "Technologies Budgie"
         or "LogDock" or any confusingly similar variation thereof;
     (d) Use the Trademarks in any manner that suggests affiliation with or
         endorsement by Technologies Budgie without written permission.

     This Section does not restrict accurate, nominative, non-commercial
     reference to the Software by its correct name in documentation, academic
     publications, or community discussions.

3.4  NO SUBLICENSING OR TRANSFER
     You may not sublicense, sell, rent, lease, transfer, assign, or otherwise
     convey any of Your rights under this Agreement to any third party. Each
     recipient of the Software receives their rights directly from the Licensor
     under the terms of this Agreement.

3.5  ARTIFICIAL INTELLIGENCE AND AUTOMATED TRAINING RESTRICTIONS
     You may not:

     (a) Use the Software, its Source Code, its internal data formats, its API
         specifications, or its algorithmic logic as training data, fine-tuning
         data, or evaluation data for any machine learning model, large language
         model, or AI system intended to compete with or replicate the
         functionality of LogDock;
     (b) Use automated means to extract, scrape, or systematically harvest
         LogDock's APIs, internal protocols, or data schemas for the purpose of
         building a competing product;
     (c) Deploy or operate LogDock in a manner where an AI system external to
         LogDock automatically reconfigures, patches, or modifies LogDock's
         operational parameters without human review, in a production environment
         serving third parties.

     For clarity: using LogDock's CE-tier AI Triage Integration for Your own
     internal log processing is fully permitted.

3.6  STORAGE BACKEND RESTRICTIONS
     You may connect LogDock CE to LVM, NAS, ZFS, and other Storage Backends
     for Internal Use and Minor Commercial Use. You may not, under this
     Agreement, offer a managed storage service or managed backup service to
     third parties where LogDock's Storage Backend integration is a primary
     component of the value delivered, without an Enterprise Edition license.

3.7  CONTAINER AND BINARY DISTRIBUTION RESTRICTIONS
     When distributing Container Deployments (e.g., Docker images) or Go
     Binaries under this Agreement:

     (a) The container image or binary must not include Enterprise Features;
     (b) The License Notice must be embedded in the image or binary metadata
         and included in accompanying documentation;
     (c) You must not strip, obfuscate, or alter the license identification
         strings embedded in the compiled binary;
     (d) Container images published to public registries must include the full
         text of this Agreement in the image's documentation layer or in the
         associated registry listing.

3.8  NO CIRCUMVENTION
     You may not use technical, contractual, or other means to circumvent any
     restriction in this Agreement. This includes, without limitation: creating
     corporate structures, subsidiaries, or shell entities to avoid the user
     count thresholds of Minor Commercial Use; using APIs, plugins, or wrappers
     to expose LogDock functionality to third parties while claiming the
     deployment is "internal"; or applying legal constructions to characterize
     a SaaS Use as Internal Use.

--- FRANÇAIS ---

Les droits accordés à la Section 2 sont soumis aux restrictions suivantes. La
violation de toute restriction de la présente Section constitue une violation
matérielle du présent Contrat et met fin automatiquement à vos droits en vertu
de la Section 11.

3.1  PAS D'USAGE SAAS SANS LICENCE ENTREPRISE
     Vous ne pouvez pas utiliser le Logiciel, ou toute Œuvre Dérivée, pour
     fournir un Service Concurrent ou pour exploiter tout Usage SaaS accessible
     à des tiers, sans avoir préalablement obtenu une licence LogDock Édition
     Entreprise auprès de Technologies Budgie. Cette restriction s'applique
     indépendamment de :

     (a) Que vous facturiez ou non le service ou que vous l'offriez gratuitement ;
     (b) Que le service soit commercialisé ou marqué sous un nom autre que
         « LogDock » ;
     (c) Que vous ayez modifié le Logiciel avant le déploiement ;
     (d) Que le Logiciel soit déployé en tant que conteneur Docker, binaire Go,
         instance en nuage ou toute autre forme ;
     (e) Que le service soit décrit comme un « enveloppeur », une « couche »,
         un « connecteur », une « intégration » ou toute autre description
         fonctionnelle qui dissimule sa dépendance à LogDock.

     Pour éviter tout doute : l'exploitation de LogDock en tant que plateforme
     interne partagée accessible aux employés et contractuels d'une seule
     organisation et de ses Affiliés directs NE constitue PAS un Usage SaaS, à
     condition qu'aucun tiers (incluant les clients, partenaires ou le grand
     public) n'ait accès aux fonctionnalités propulsées par LogDock.

3.2  PAS D'EXTRACTION OU DISTRIBUTION DE FONCTIONNALITÉS ENTREPRISE
     Vous ne pouvez pas, dans le cadre du présent Contrat :

     (a) Extraire, copier, porter, réimplémenter ou distribuer toute
         Fonctionnalité Entreprise ou tout code implémentant une Fonctionnalité
         Entreprise ;
     (b) Modifier le Logiciel pour supprimer les mécanismes de contrôle d'accès
         qui restreignent l'accès aux Fonctionnalités Entreprise, que ces
         mécanismes soient implémentés via des vérifications de clé de licence,
         des balises de compilation, des indicateurs de compilation, des
         vérifications d'exécution ou tout autre moyen technique ;
     (c) Distribuer toute version du Logiciel dans laquelle les Fonctionnalités
         Entreprise ont été déverrouillées, activées ou rendues accessibles sans
         une licence Édition Entreprise valide ;
     (d) Procéder à l'ingénierie inverse des Fonctionnalités Entreprise dans le
         but de les réimplémenter dans un produit concurrent, que vous ayez ou
         non accès au code source correspondant.

3.3  RESTRICTIONS RELATIVES AUX MARQUES DE COMMERCE
     Vous ne pouvez pas :

     (a) Utiliser les Marques de Commerce « Technologies Budgie » ou « LogDock »
         dans le nom, la marque, le marketing ou la documentation de tout
         produit, service ou organisation qui n'est pas Technologies Budgie,
         sans le consentement écrit préalable de Technologies Budgie ;
     (b) Représenter que votre Œuvre Dérivée ou service est le produit LogDock
         officiel ou est approuvé, parrainé ou soutenu par Technologies Budgie ;
     (c) Enregistrer ou tenter d'enregistrer toute marque de commerce, nom de
         domaine, identifiant de médias sociaux ou nom commercial qui incorpore
         « Technologies Budgie » ou « LogDock » ou toute variation prêtant à
         confusion ;
     (d) Utiliser les Marques de Commerce de quelque manière que ce soit qui
         suggère une affiliation avec ou une approbation de Technologies Budgie
         sans autorisation écrite.

     La présente Section ne restreint pas la référence nominative exacte et non
     commerciale au Logiciel par son nom correct dans la documentation, les
     publications académiques ou les discussions communautaires.

3.4  PAS DE SOUS-LICENCE OU DE TRANSFERT
     Vous ne pouvez pas sous-licencier, vendre, louer, céder à bail, transférer,
     céder ou autrement transmettre l'un ou l'autre de vos droits dans le cadre
     du présent Contrat à un tiers. Chaque destinataire du Logiciel reçoit ses
     droits directement du Concédant dans le cadre du présent Contrat.

3.5  RESTRICTIONS RELATIVES À L'INTELLIGENCE ARTIFICIELLE ET À L'AUTOMATISATION
     Vous ne pouvez pas :

     (a) Utiliser le Logiciel, son Code Source, ses formats de données internes,
         ses spécifications d'API ou sa logique algorithmique comme données
         d'entraînement, données d'affinage ou données d'évaluation pour tout
         modèle d'apprentissage automatique, grand modèle de langage ou système
         d'IA destiné à concurrencer ou reproduire les fonctionnalités de
         LogDock ;
     (b) Utiliser des moyens automatisés pour extraire, aspirer ou récolter
         systématiquement les API, protocoles internes ou schémas de données de
         LogDock dans le but de construire un produit concurrent ;
     (c) Déployer ou exploiter LogDock d'une manière où un système d'IA externe
         à LogDock reconfigure, corrige ou modifie automatiquement les paramètres
         opérationnels de LogDock sans examen humain, dans un environnement de
         production servant des tiers.

     Pour plus de clarté : l'utilisation de l'Intégration IA Triage de niveau
     EC de LogDock pour votre propre traitement interne des journaux est
     entièrement permise.

3.6  RESTRICTIONS RELATIVES AUX BACKENDS DE STOCKAGE
     Vous pouvez connecter LogDock EC aux Backends de Stockage LVM, NAS, ZFS et
     autres pour l'Usage Interne et l'Usage Commercial Mineur. Vous ne pouvez
     pas, dans le cadre du présent Contrat, offrir un service de stockage géré
     ou un service de sauvegarde géré à des tiers où l'intégration du Backend de
     Stockage de LogDock est une composante principale de la valeur fournie, sans
     une licence Édition Entreprise.

3.7  RESTRICTIONS RELATIVES À LA DISTRIBUTION DE CONTENEURS ET BINAIRES
     Lors de la distribution de Déploiements en Conteneur (p. ex., images
     Docker) ou de Binaires Go dans le cadre du présent Contrat :

     (a) L'image de conteneur ou le binaire ne doit pas inclure de
         Fonctionnalités Entreprise ;
     (b) L'Avis de Licence doit être intégré dans les métadonnées de l'image ou
         du binaire et inclus dans la documentation d'accompagnement ;
     (c) Vous ne devez pas supprimer, obscurcir ou modifier les chaînes
         d'identification de licence intégrées dans le binaire compilé ;
     (d) Les images de conteneurs publiées dans des registres publics doivent
         inclure le texte intégral du présent Contrat dans la couche de
         documentation de l'image ou dans la liste du registre associé.

3.8  PAS DE CONTOURNEMENT
     Vous ne pouvez pas utiliser des moyens techniques, contractuels ou autres
     pour contourner toute restriction du présent Contrat. Cela inclut, sans
     s'y limiter : la création de structures corporatives, de filiales ou
     d'entités fictives pour éviter les seuils de comptage d'utilisateurs de
     l'Usage Commercial Mineur ; l'utilisation d'API, de plugins ou
     d'enveloppeurs pour exposer les fonctionnalités de LogDock à des tiers tout
     en prétendant que le déploiement est « interne » ; ou l'application de
     constructions juridiques pour caractériser un Usage SaaS comme un Usage
     Interne.

================================================================================
SECTION 4 — LICENSE NOTICE REQUIREMENTS / EXIGENCES D'AVIS DE LICENCE
================================================================================

--- ENGLISH ---

4.1  REQUIRED NOTICE TEXT
     All copies, Distributions, and Derivative Works of the Software must
     include the following notice, preserved in full and unmodified:

     -------------------------------------------------------------------------
     LogDock Community Edition
     Copyright (C) 2025 Technologies Budgie. All rights reserved.
     Website: https://technologiesbudgie.page.dev
     Contact: budgie@mailfence.com

     This software is licensed under the LogDock Community Edition License
     Agreement (LDCEL) v1.0. You may use, copy, modify, and distribute this
     software subject to the terms of the LDCEL. Commercial SaaS use,
     distribution of Enterprise Features, and creation of competing hosted
     services are prohibited without an Enterprise Edition license.

     Full license text: https://technologiesbudgie.page.dev/licenses/LDCEL
     -------------------------------------------------------------------------

4.2  PLACEMENT OF NOTICE
     The License Notice must appear:

     (a) In every Source Code file modified or distributed by You, at the top
         of the file or in a clearly visible comment block;
     (b) In a file named LICENSE, LICENSE.txt, or COPYING in the root of any
         source code repository or archive;
     (c) In the documentation or README of any distributed package;
     (d) In the metadata of any Docker image (via LABEL instructions);
     (e) In the --version or --license output of any distributed binary;
     (f) On any website or service page from which the Software or Derivative
         Work is made available for download.

4.3  PRESERVATION OF NOTICES
     You must not remove, alter, obscure, or supplement any copyright notice,
     Trademark attribution, or License Notice contained in the original
     Software. If You add Your own copyright notices (for Your original
     contributions), You must do so in a manner that does not suggest
     Technologies Budgie's endorsement of or responsibility for Your
     contributions.

--- FRANÇAIS ---

4.1  TEXTE D'AVIS REQUIS
     Toutes les copies, Distributions et Œuvres Dérivées du Logiciel doivent
     inclure l'avis suivant, conservé dans son intégralité et non modifié :

     -------------------------------------------------------------------------
     LogDock Édition Communautaire
     Copyright (C) 2025 Technologies Budgie. Tous droits réservés.
     Site web : https://technologiesbudgie.page.dev
     Contact : budgie@mailfence.com

     Ce logiciel est licencié sous le Contrat de Licence LogDock Édition
     Communautaire (LDLEC) v1.0. Vous pouvez utiliser, copier, modifier et
     distribuer ce logiciel sous réserve des conditions du LDLEC. L'usage SaaS
     commercial, la distribution des Fonctionnalités Entreprise et la création
     de services hébergés concurrents sont interdits sans une licence Édition
     Entreprise.

     Texte complet de la licence : https://technologiesbudgie.page.dev/licenses/LDLEC
     -------------------------------------------------------------------------

4.2  EMPLACEMENT DE L'AVIS
     L'Avis de Licence doit apparaître :

     (a) Dans chaque fichier de Code Source modifié ou distribué par Vous, en
         haut du fichier ou dans un bloc de commentaires clairement visible ;
     (b) Dans un fichier nommé LICENSE, LICENSE.txt ou COPYING à la racine de
         tout dépôt ou archive de code source ;
     (c) Dans la documentation ou le README de tout paquet distribué ;
     (d) Dans les métadonnées de toute image Docker (via des instructions LABEL) ;
     (e) Dans la sortie --version ou --license de tout binaire distribué ;
     (f) Sur tout site web ou page de service à partir duquel le Logiciel ou
         l'Œuvre Dérivée est mis à disposition pour téléchargement.

4.3  PRÉSERVATION DES AVIS
     Vous ne devez pas supprimer, modifier, obscurcir ou compléter tout avis de
     droit d'auteur, attribution de Marque de Commerce ou Avis de Licence
     contenu dans le Logiciel original. Si vous ajoutez vos propres avis de
     droit d'auteur (pour vos contributions originales), vous devez le faire
     d'une manière qui ne suggère pas l'approbation de Technologies Budgie ou
     sa responsabilité à l'égard de vos contributions.

================================================================================
SECTION 5 — REDISTRIBUTION RULES / RÈGLES DE REDISTRIBUTION
================================================================================

--- ENGLISH ---

5.1  COPYLEFT OBLIGATION
     If You Distribute the Software or any Derivative Work in Source Code form,
     You must license the entire Derivative Work under this Agreement. You may
     not impose any additional restrictions on the exercise of rights granted
     under this Agreement beyond those already contained herein.

5.2  NETWORK USE COPYLEFT
     If You modify the Software and operate the modified version to provide
     services over a network (including internal networks), You must make the
     complete corresponding Source Code of Your modified version available to
     the users of that service under the terms of this Agreement. You must
     provide a prominent notice in the user interface or documentation of the
     service explaining how to obtain the Source Code.

5.3  COMPATIBLE LICENSES
     You may not distribute the Software under any license other than this
     Agreement. Combinations with other software are permitted provided the
     other software components are governed by their own licenses and the
     LogDock components remain governed by this Agreement without additional
     restrictions.

5.4  PACKAGE REGISTRY DISTRIBUTIONS
     If You publish the Software or a Derivative Work to any public package
     registry (including but not limited to Docker Hub, GitHub Container
     Registry, npm, PyPI, or any Go module proxy), You must:

     (a) Set the license field to "LicenseRef-LDCEL-1.0" or an equivalent
         unambiguous identifier;
     (b) Include a link to the full license text;
     (c) Ensure the listing prominently notes that SaaS use requires a separate
         Enterprise Edition license.

5.5  NO TIVOIZATION
     If You Distribute the Software in Object Code or binary form on a device or
     platform that uses technical measures to prevent users from installing or
     running modified versions of the Software on that device, You must provide
     the installation information necessary for users to do so, to the extent
     such information exists and You are permitted to provide it.

--- FRANÇAIS ---

5.1  OBLIGATION COPYLEFT
     Si vous Distribuez le Logiciel ou toute Œuvre Dérivée sous forme de Code
     Source, vous devez licencier l'ensemble de l'Œuvre Dérivée dans le cadre
     du présent Contrat. Vous ne pouvez pas imposer de restrictions
     supplémentaires à l'exercice des droits accordés dans le cadre du présent
     Contrat au-delà de ceux déjà contenus dans les présentes.

5.2  COPYLEFT D'USAGE RÉSEAU
     Si vous modifiez le Logiciel et exploitez la version modifiée pour fournir
     des services sur un réseau (incluant les réseaux internes), vous devez
     mettre à disposition le Code Source correspondant complet de votre version
     modifiée aux utilisateurs de ce service dans le cadre du présent Contrat.
     Vous devez fournir un avis bien visible dans l'interface utilisateur ou la
     documentation du service expliquant comment obtenir le Code Source.

5.3  LICENCES COMPATIBLES
     Vous ne pouvez pas distribuer le Logiciel sous une licence autre que le
     présent Contrat. Les combinaisons avec d'autres logiciels sont permises à
     condition que les autres composants logiciels soient régis par leurs propres
     licences et que les composants LogDock restent régis par le présent Contrat
     sans restrictions supplémentaires.

5.4  DISTRIBUTIONS DANS DES REGISTRES DE PAQUETS
     Si vous publiez le Logiciel ou une Œuvre Dérivée dans tout registre de
     paquets public (incluant notamment Docker Hub, GitHub Container Registry,
     npm, PyPI ou tout proxy de module Go), vous devez :

     (a) Définir le champ de licence sur « LicenseRef-LDCEL-1.0 » ou un
         identifiant non ambigu équivalent ;
     (b) Inclure un lien vers le texte complet de la licence ;
     (c) S'assurer que la liste indique bien que l'usage SaaS nécessite une
         licence Édition Entreprise séparée.

5.5  PAS DE TIVOISATON
     Si vous Distribuez le Logiciel sous forme de Code Objet ou de binaire sur
     un appareil ou une plateforme qui utilise des mesures techniques pour
     empêcher les utilisateurs d'installer ou d'exécuter des versions modifiées
     du Logiciel sur cet appareil, vous devez fournir les informations
     d'installation nécessaires pour que les utilisateurs puissent le faire,
     dans la mesure où ces informations existent et où vous êtes autorisé à les
     fournir.

================================================================================
SECTION 6 — CONTRIBUTOR TERMS / CONDITIONS DES CONTRIBUTEURS
================================================================================

--- ENGLISH ---

6.1  CONTRIBUTOR LICENSE GRANT
     By submitting a Contribution to Technologies Budgie (whether via pull
     request, patch, email, issue tracker, or any other means), You hereby
     grant to Technologies Budgie a perpetual, worldwide, non-exclusive,
     royalty-free, irrevocable license to:

     (a) Reproduce, prepare Derivative Works of, publicly display, publicly
         perform, sublicense (including under commercial licenses), and
         distribute Your Contribution and such Derivative Works;
     (b) Include Your Contribution in both the Community Edition and the
         Enterprise Edition of LogDock;
     (c) Re-license Your Contribution under the LogDock Enterprise Edition
         License Agreement or any future license adopted by Technologies Budgie
         for the Software.

6.2  PATENT LICENSE FROM CONTRIBUTORS
     Subject to the terms of this Agreement, each Contributor hereby grants to
     Technologies Budgie and to all recipients of the Software a perpetual,
     worldwide, non-exclusive, royalty-free, irrevocable patent license to make,
     have made, use, offer to sell, sell, import, and otherwise transfer the
     Software to the extent such license applies to Your Contribution.

6.3  REPRESENTATIONS BY CONTRIBUTORS
     By submitting a Contribution, You represent that:

     (a) You have the legal right to grant the licenses described in this
         Section;
     (b) The Contribution is Your original work or You have obtained all
         necessary rights and permissions from third parties;
     (c) The Contribution does not knowingly infringe any third-party copyright,
         patent, trademark, or other intellectual property right;
     (d) The Contribution does not include any code, content, or material that
         is subject to terms that would impose obligations on Technologies Budgie
         beyond those of this Agreement.

6.4  NO OBLIGATION TO ACCEPT
     Technologies Budgie is under no obligation to accept, incorporate, or
     maintain any Contribution. Technologies Budgie may reject, modify, or
     replace any Contribution at its sole discretion.

6.5  CONTRIBUTOR ATTRIBUTION
     Technologies Budgie will maintain a record of Contributors in the project's
     CONTRIBUTORS file or equivalent. Attribution does not grant any right to
     control or influence the development direction of LogDock.

6.6  CORPORATE CONTRIBUTORS
     If You submit a Contribution on behalf of a legal entity, You represent
     that You are authorized to bind that entity to the terms of this Section.

--- FRANÇAIS ---

6.1  OCTROI DE LICENCE DU CONTRIBUTEUR
     En soumettant une Contribution à Technologies Budgie (que ce soit via une
     demande de tirage, un correctif, un courriel, un système de suivi des
     problèmes ou tout autre moyen), vous accordez par les présentes à
     Technologies Budgie une licence perpétuelle, mondiale, non exclusive, libre
     de redevances et irrévocable pour :

     (a) Reproduire, préparer des Œuvres Dérivées de, afficher publiquement,
         exécuter publiquement, sous-licencier (incluant sous des licences
         commerciales) et distribuer votre Contribution et ces Œuvres Dérivées ;
     (b) Inclure votre Contribution dans l'Édition Communautaire et l'Édition
         Entreprise de LogDock ;
     (c) Relicencier votre Contribution sous le Contrat de Licence LogDock
         Édition Entreprise ou toute licence future adoptée par Technologies
         Budgie pour le Logiciel.

6.2  LICENCE DE BREVET DES CONTRIBUTEURS
     Sous réserve des conditions du présent Contrat, chaque Contributeur accorde
     par les présentes à Technologies Budgie et à tous les destinataires du
     Logiciel une licence de brevet perpétuelle, mondiale, non exclusive, libre
     de redevances et irrévocable pour fabriquer, faire fabriquer, utiliser,
     offrir à la vente, vendre, importer et autrement transférer le Logiciel
     dans la mesure où cette licence s'applique à votre Contribution.

6.3  DÉCLARATIONS DES CONTRIBUTEURS
     En soumettant une Contribution, vous déclarez que :

     (a) Vous avez le droit légal d'accorder les licences décrites dans la
         présente Section ;
     (b) La Contribution est votre œuvre originale ou vous avez obtenu tous les
         droits et permissions nécessaires de tiers ;
     (c) La Contribution ne porte pas atteinte, à votre connaissance, à tout
         droit d'auteur, brevet, marque de commerce ou autre droit de propriété
         intellectuelle de tiers ;
     (d) La Contribution n'inclut aucun code, contenu ou matériel soumis à des
         conditions qui imposeraient des obligations à Technologies Budgie
         au-delà de celles du présent Contrat.

6.4  AUCUNE OBLIGATION D'ACCEPTATION
     Technologies Budgie n'a aucune obligation d'accepter, d'incorporer ou de
     maintenir toute Contribution. Technologies Budgie peut rejeter, modifier ou
     remplacer toute Contribution à sa seule discrétion.

6.5  ATTRIBUTION DES CONTRIBUTEURS
     Technologies Budgie maintiendra un registre des Contributeurs dans le
     fichier CONTRIBUTORS du projet ou équivalent. L'attribution n'accorde
     aucun droit de contrôler ou d'influencer la direction du développement de
     LogDock.

6.6  CONTRIBUTEURS CORPORATIFS
     Si vous soumettez une Contribution au nom d'une entité juridique, vous
     déclarez être autorisé à lier cette entité aux conditions de la présente
     Section.

================================================================================
SECTION 7 — WARRANTIES AND DISCLAIMER / GARANTIES ET AVIS DE NON-RESPONSABILITÉ
================================================================================

--- ENGLISH ---

7.1  NO WARRANTY
     THE SOFTWARE IS PROVIDED "AS IS," WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
     IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
     FITNESS FOR A PARTICULAR PURPOSE, TITLE, ACCURACY, RELIABILITY,
     COMPLETENESS, AND NON-INFRINGEMENT. TECHNOLOGIES BUDGIE DOES NOT WARRANT
     THAT THE SOFTWARE WILL MEET YOUR REQUIREMENTS, THAT OPERATION OF THE
     SOFTWARE WILL BE UNINTERRUPTED OR ERROR-FREE, THAT DEFECTS WILL BE
     CORRECTED, OR THAT THE SOFTWARE IS FREE OF VIRUSES OR OTHER HARMFUL
     COMPONENTS.

7.2  STORAGE BACKEND DISCLAIMER
     TECHNOLOGIES BUDGIE EXPRESSLY DISCLAIMS ALL WARRANTIES WITH RESPECT TO THE
     OPERATION OF LOGDOCK IN CONJUNCTION WITH ANY STORAGE BACKEND, INCLUDING
     LVM, NAS, AND ZFS SYSTEMS. THE LICENSEE ASSUMES ALL RISK ASSOCIATED WITH
     DATA STORAGE, DATA LOSS, DATA CORRUPTION, FILESYSTEM DAMAGE, VOLUME
     MANAGEMENT ERRORS, AND ANY OTHER CONSEQUENCES ARISING FROM THE INTEGRATION
     OF LOGDOCK WITH STORAGE BACKENDS.

7.3  AI TRIAGE DISCLAIMER
     THE AI TRIAGE INTEGRATION FEATURES OF THE SOFTWARE ARE PROVIDED FOR
     INFORMATIONAL AND OPERATIONAL ASSISTANCE PURPOSES ONLY. TECHNOLOGIES
     BUDGIE DOES NOT WARRANT THAT AI TRIAGE OUTPUTS ARE ACCURATE, COMPLETE,
     OR SUITABLE FOR ANY PARTICULAR PURPOSE. THE LICENSEE ASSUMES FULL
     RESPONSIBILITY FOR ANY DECISIONS MADE ON THE BASIS OF AI TRIAGE OUTPUTS.

7.4  QUÉBEC LAW LIMITATION
     CERTAIN JURISDICTIONS, INCLUDING THE PROVINCE OF QUÉBEC UNDER THE CONSUMER
     PROTECTION ACT (RLRQ C P-40.1), DO NOT ALLOW THE EXCLUSION OF IMPLIED
     WARRANTIES IN CERTAIN CIRCUMSTANCES. TO THE EXTENT THAT SUCH LAWS APPLY
     AND CANNOT BE WAIVED BY AGREEMENT, THE EXCLUSIONS SET FORTH ABOVE ARE
     MODIFIED TO THE MINIMUM EXTENT NECESSARY TO COMPLY WITH SUCH LAWS.
     THIS SECTION 7.4 APPLIES ONLY WHERE REQUIRED BY MANDATORY LAW.

--- FRANÇAIS ---

7.1  AUCUNE GARANTIE
     LE LOGICIEL EST FOURNI « TEL QUEL », SANS GARANTIE D'AUCUNE SORTE,
     EXPRESSE OU IMPLICITE, INCLUANT NOTAMMENT LES GARANTIES DE QUALITÉ
     MARCHANDE, D'ADÉQUATION À UN USAGE PARTICULIER, DE TITRE, D'EXACTITUDE,
     DE FIABILITÉ, D'EXHAUSTIVITÉ ET DE NON-VIOLATION. TECHNOLOGIES BUDGIE NE
     GARANTIT PAS QUE LE LOGICIEL RÉPONDRA À VOS EXIGENCES, QUE L'EXPLOITATION
     DU LOGICIEL SERA ININTERROMPUE OU SANS ERREUR, QUE LES DÉFAUTS SERONT
     CORRIGÉS, OU QUE LE LOGICIEL EST EXEMPT DE VIRUS OU D'AUTRES COMPOSANTS
     NUISIBLES.

7.2  AVIS DE NON-RESPONSABILITÉ RELATIF AUX BACKENDS DE STOCKAGE
     TECHNOLOGIES BUDGIE DÉCLINE EXPRESSÉMENT TOUTE GARANTIE CONCERNANT LE
     FONCTIONNEMENT DE LOGDOCK EN CONJONCTION AVEC TOUT BACKEND DE STOCKAGE,
     INCLUANT LES SYSTÈMES LVM, NAS ET ZFS. LE LICENCIÉ ASSUME TOUS LES RISQUES
     ASSOCIÉS AU STOCKAGE DE DONNÉES, À LA PERTE DE DONNÉES, À LA CORRUPTION
     DES DONNÉES, AUX DOMMAGES AU SYSTÈME DE FICHIERS, AUX ERREURS DE GESTION
     DES VOLUMES ET À TOUTE AUTRE CONSÉQUENCE DÉCOULANT DE L'INTÉGRATION DE
     LOGDOCK AVEC LES BACKENDS DE STOCKAGE.

7.3  AVIS DE NON-RESPONSABILITÉ RELATIF AU TRIAGE IA
     LES FONCTIONNALITÉS D'INTÉGRATION IA TRIAGE DU LOGICIEL SONT FOURNIES À
     DES FINS D'ASSISTANCE INFORMATIONNELLE ET OPÉRATIONNELLE UNIQUEMENT.
     TECHNOLOGIES BUDGIE NE GARANTIT PAS QUE LES RÉSULTATS DU TRIAGE IA SONT
     EXACTS, COMPLETS OU ADAPTÉS À UN USAGE PARTICULIER. LE LICENCIÉ ASSUME
     L'ENTIÈRE RESPONSABILITÉ DE TOUTE DÉCISION PRISE SUR LA BASE DES RÉSULTATS
     DU TRIAGE IA.

7.4  LIMITATION EN VERTU DU DROIT QUÉBÉCOIS
     CERTAINES JURIDICTIONS, INCLUANT LA PROVINCE DE QUÉBEC EN VERTU DE LA LOI
     SUR LA PROTECTION DU CONSOMMATEUR (RLRQ C P-40.1), NE PERMETTENT PAS
     L'EXCLUSION DES GARANTIES IMPLICITES DANS CERTAINES CIRCONSTANCES. DANS LA
     MESURE OÙ DE TELLES LOIS S'APPLIQUENT ET NE PEUVENT ÊTRE RENONCÉES PAR
     ACCORD, LES EXCLUSIONS ÉNONCÉES CI-DESSUS SONT MODIFIÉES DANS LA MESURE
     MINIMALE NÉCESSAIRE POUR SE CONFORMER À DE TELLES LOIS. LA PRÉSENTE
     SECTION 7.4 S'APPLIQUE UNIQUEMENT LÀ OÙ LA LOI IMPÉRATIVE L'EXIGE.

================================================================================
SECTION 8 — LIMITATION OF LIABILITY / LIMITATION DE RESPONSABILITÉ
================================================================================

--- ENGLISH ---

8.1  EXCLUSION OF CONSEQUENTIAL DAMAGES
     TO THE MAXIMUM EXTENT PERMITTED BY APPLICABLE LAW, IN NO EVENT SHALL
     TECHNOLOGIES BUDGIE, ITS DIRECTORS, OFFICERS, EMPLOYEES, AGENTS,
     CONTRACTORS, OR CONTRIBUTORS BE LIABLE FOR ANY INDIRECT, INCIDENTAL,
     SPECIAL, EXEMPLARY, PUNITIVE, OR CONSEQUENTIAL DAMAGES WHATSOEVER,
     INCLUDING BUT NOT LIMITED TO: LOSS OF PROFITS, LOSS OF REVENUE, LOSS OF
     DATA, LOSS OF BUSINESS, BUSINESS INTERRUPTION, LOSS OF GOODWILL, COST OF
     SUBSTITUTE GOODS OR SERVICES, OR ANY OTHER PECUNIARY LOSS, EVEN IF
     TECHNOLOGIES BUDGIE HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGES,
     AND WHETHER ARISING UNDER ANY THEORY OF LIABILITY, INCLUDING CONTRACT,
     TORT (INCLUDING NEGLIGENCE), STRICT LIABILITY, OR OTHERWISE.

8.2  CAP ON LIABILITY
     TO THE MAXIMUM EXTENT PERMITTED BY APPLICABLE LAW, TECHNOLOGIES BUDGIE'S
     TOTAL CUMULATIVE LIABILITY TO YOU ARISING OUT OF OR RELATING TO THIS
     AGREEMENT OR THE SOFTWARE, REGARDLESS OF THE FORM OR NATURE OF THE CLAIM,
     SHALL NOT EXCEED ONE HUNDRED CANADIAN DOLLARS (CAD $100.00). THIS CAP
     APPLIES IN AGGREGATE ACROSS ALL CLAIMS AND CAUSES OF ACTION.

8.3  ESSENTIAL BASIS
     YOU ACKNOWLEDGE THAT TECHNOLOGIES BUDGIE HAS SET ITS PRICING (INCLUDING
     PROVIDING THE COMMUNITY EDITION AT NO CHARGE) IN RELIANCE UPON THE
     LIMITATIONS OF LIABILITY SET FORTH HEREIN, AND THAT SUCH LIMITATIONS FORM
     AN ESSENTIAL BASIS OF THE BARGAIN BETWEEN THE PARTIES. THE LIMITATIONS
     SHALL APPLY NOTWITHSTANDING ANY FAILURE OF ESSENTIAL PURPOSE OF ANY LIMITED
     REMEDY.

8.4  INDEMNIFICATION BY LICENSEE
     You agree to indemnify, defend, and hold harmless Technologies Budgie, its
     directors, officers, employees, agents, and Contributors from and against
     any and all claims, liabilities, damages, losses, penalties, fines, costs,
     and expenses (including reasonable legal fees) arising out of or relating
     to: (a) Your use of the Software; (b) Your violation of this Agreement;
     (c) Your operation of any service using the Software; (d) any Derivative
     Work created by You; or (e) any claim by a third party arising from Your
     Distribution of the Software.

--- FRANÇAIS ---

8.1  EXCLUSION DES DOMMAGES INDIRECTS
     DANS LA MESURE MAXIMALE PERMISE PAR LA LOI APPLICABLE, EN AUCUN CAS
     TECHNOLOGIES BUDGIE, SES ADMINISTRATEURS, DIRIGEANTS, EMPLOYÉS, AGENTS,
     CONTRACTUELS OU CONTRIBUTEURS NE SERONT RESPONSABLES DE TOUT DOMMAGE
     INDIRECT, ACCESSOIRE, SPÉCIAL, EXEMPLAIRE, PUNITIF OU CONSÉCUTIF, INCLUANT
     NOTAMMENT : LA PERTE DE PROFITS, LA PERTE DE REVENUS, LA PERTE DE DONNÉES,
     LA PERTE D'ENTREPRISE, L'INTERRUPTION DES ACTIVITÉS, LA PERTE DE
     CLIENTÈLE, LE COÛT DE BIENS OU SERVICES DE REMPLACEMENT, OU TOUTE AUTRE
     PERTE PÉCUNIAIRE, MÊME SI TECHNOLOGIES BUDGIE A ÉTÉ AVISÉ DE LA
     POSSIBILITÉ DE TELS DOMMAGES, ET QU'ILS SURVIENNENT DANS LE CADRE DE TOUTE
     THÉORIE DE RESPONSABILITÉ, INCLUANT LE CONTRAT, LA RESPONSABILITÉ CIVILE
     (INCLUANT LA NÉGLIGENCE), LA RESPONSABILITÉ STRICTE OU AUTRE.

8.2  PLAFOND DE RESPONSABILITÉ
     DANS LA MESURE MAXIMALE PERMISE PAR LA LOI APPLICABLE, LA RESPONSABILITÉ
     CUMULATIVE TOTALE DE TECHNOLOGIES BUDGIE ENVERS VOUS DÉCOULANT DU PRÉSENT
     CONTRAT OU DU LOGICIEL, OU S'Y RAPPORTANT, QUELLE QUE SOIT LA FORME OU LA
     NATURE DE LA RÉCLAMATION, NE DÉPASSERA PAS CENT DOLLARS CANADIENS (100 $
     CAD). CE PLAFOND S'APPLIQUE DE MANIÈRE GLOBALE À TOUTES LES RÉCLAMATIONS ET
     CAUSES D'ACTION.

8.3  BASE ESSENTIELLE
     VOUS RECONNAISSEZ QUE TECHNOLOGIES BUDGIE A ÉTABLI SES PRIX (INCLUANT LA
     FOURNITURE GRATUITE DE L'ÉDITION COMMUNAUTAIRE) EN SE FONDANT SUR LES
     LIMITATIONS DE RESPONSABILITÉ ÉNONCÉES AUX PRÉSENTES, ET QUE CES
     LIMITATIONS CONSTITUENT UNE BASE ESSENTIELLE DE L'ACCORD ENTRE LES
     PARTIES. LES LIMITATIONS S'APPLIQUERONT NONOBSTANT TOUT DÉFAUT DE L'OBJET
     ESSENTIEL DE TOUT RECOURS LIMITÉ.

8.4  INDEMNISATION PAR LE LICENCIÉ
     Vous acceptez d'indemniser, de défendre et de dégager de toute
     responsabilité Technologies Budgie, ses administrateurs, dirigeants,
     employés, agents et Contributeurs contre toute réclamation, responsabilité,
     dommage, perte, pénalité, amende, coût et dépense (incluant les honoraires
     juridiques raisonnables) découlant de ou liés à : (a) Votre utilisation du
     Logiciel ; (b) Votre violation du présent Contrat ; (c) Votre exploitation
     de tout service utilisant le Logiciel ; (d) toute Œuvre Dérivée créée par
     Vous ; ou (e) toute réclamation d'un tiers découlant de Votre Distribution
     du Logiciel.

================================================================================
SECTION 9 — DUAL LICENSING / DOUBLE LICENCE
================================================================================

--- ENGLISH ---

9.1  UPGRADE PATH TO ENTERPRISE EDITION
     This Agreement governs only the Community Edition of LogDock. If You
     require rights beyond those granted herein — including the right to operate
     LogDock as a SaaS or managed service for third parties, the right to use
     Enterprise Features, the right to deploy LogDock in multi-tenant commercial
     environments, or the right to remove copyleft obligations — You must obtain
     a separate LogDock Enterprise Edition license.

9.2  HOW TO OBTAIN AN ENTERPRISE LICENSE
     To obtain a LogDock Enterprise Edition license, contact Technologies Budgie
     at:

         Email  : budgie@mailfence.com
         Website: https://technologiesbudgie.page.dev

     Enterprise licenses are negotiated on a case-by-case basis and may include
     per-seat, per-node, per-deployment, revenue-based, or other pricing models
     as agreed between the parties.

9.3  COMMUNITY TO ENTERPRISE TRANSITION
     Organizations currently operating under this Community Edition License that
     wish to transition to the Enterprise Edition must:

     (a) Contact Technologies Budgie at the address above;
     (b) Execute the LogDock Enterprise Edition License Agreement;
     (c) Pay applicable Enterprise Edition fees as agreed;
     (d) Upon execution of the Enterprise Edition License, the Enterprise Edition
         License shall govern all subsequent use of LogDock by that organization,
         and this Agreement shall no longer apply to that organization's use.

9.4  NO AUTOMATIC UPGRADE
     Acceptance of a Contribution by Technologies Budgie or participation in the
     LogDock community does not automatically grant Enterprise Edition rights.
     Enterprise rights are only obtained through a signed commercial agreement.

9.5  DUAL LICENSING STATEMENT
     Technologies Budgie offers LogDock under this dual-licensing model to
     support community innovation while sustaining commercial viability. Revenue
     from Enterprise Edition licenses funds continued development of both
     editions. Technologies Budgie is committed to maintaining a feature-rich
     Community Edition as a long-term product offering.

--- FRANÇAIS ---

9.1  VOIE DE MISE À NIVEAU VERS L'ÉDITION ENTREPRISE
     Le présent Contrat régit uniquement l'Édition Communautaire de LogDock. Si
     vous nécessitez des droits au-delà de ceux accordés aux présentes —
     incluant le droit d'exploiter LogDock en tant que SaaS ou service géré pour
     des tiers, le droit d'utiliser les Fonctionnalités Entreprise, le droit de
     déployer LogDock dans des environnements commerciaux multi-locataires, ou
     le droit de supprimer les obligations copyleft — vous devez obtenir une
     licence LogDock Édition Entreprise séparée.

9.2  COMMENT OBTENIR UNE LICENCE ENTREPRISE
     Pour obtenir une licence LogDock Édition Entreprise, contactez Technologies
     Budgie à :

         Courriel : budgie@mailfence.com
         Site web : https://technologiesbudgie.page.dev

     Les licences entreprise sont négociées au cas par cas et peuvent inclure
     des modèles de tarification par siège, par nœud, par déploiement, basés sur
     les revenus ou autres selon l'accord entre les parties.

9.3  TRANSITION DE L'ÉDITION COMMUNAUTAIRE À L'ENTREPRISE
     Les organisations exploitant actuellement sous la présente Licence Édition
     Communautaire qui souhaitent passer à l'Édition Entreprise doivent :

     (a) Contacter Technologies Budgie à l'adresse ci-dessus ;
     (b) Exécuter le Contrat de Licence LogDock Édition Entreprise ;
     (c) Payer les frais applicables de l'Édition Entreprise selon l'accord ;
     (d) Lors de l'exécution de la Licence Édition Entreprise, celle-ci régira
         toute utilisation ultérieure de LogDock par cette organisation, et le
         présent Contrat ne s'appliquera plus à l'utilisation de cette
         organisation.

9.4  PAS DE MISE À NIVEAU AUTOMATIQUE
     L'acceptation d'une Contribution par Technologies Budgie ou la
     participation à la communauté LogDock n'accorde pas automatiquement des
     droits d'Édition Entreprise. Les droits d'entreprise ne sont obtenus que
     par le biais d'un accord commercial signé.

9.5  DÉCLARATION DE DOUBLE LICENCE
     Technologies Budgie offre LogDock dans le cadre de ce modèle de double
     licence pour soutenir l'innovation communautaire tout en maintenant la
     viabilité commerciale. Les revenus des licences Édition Entreprise financent
     le développement continu des deux éditions. Technologies Budgie s'engage à
     maintenir une Édition Communautaire riche en fonctionnalités en tant
     qu'offre de produit à long terme.

================================================================================
SECTION 10 — PATENTS / BREVETS
================================================================================

--- ENGLISH ---

10.1 PATENT GRANT
     Subject to the terms of this Agreement, each Contributor hereby grants You
     a non-exclusive, worldwide, royalty-free patent license under that
     Contributor's essential patent claims to make, use, sell, offer to sell,
     import, and otherwise run, modify, and propagate the contents of this
     Software, to the extent permitted by the rights granted herein.

10.2 PATENT RETALIATION
     If You institute patent litigation (including a cross-claim or counterclaim)
     against any entity alleging that the Software, or a Contribution
     incorporated within the Software, constitutes direct or contributory patent
     infringement, then all patent licenses granted to You under this Agreement
     for that Software shall terminate as of the date such litigation is filed.

10.3 TECHNOLOGIES BUDGIE PATENTS
     Nothing in this Agreement grants You any rights under Technologies Budgie's
     patents, patent applications, or trade secrets, other than the limited
     patent license described in Section 10.1 with respect to Contributions.
     Technologies Budgie's proprietary algorithms, compression methods, AI
     Triage Integration logic, and DuckDB optimization techniques may be subject
     to patent protection. No license to such patents is granted under this
     Agreement except as necessary to exercise the rights explicitly granted in
     Section 2.

--- FRANÇAIS ---

10.1 OCTROI DE BREVET
     Sous réserve des conditions du présent Contrat, chaque Contributeur vous
     accorde par les présentes une licence de brevet non exclusive, mondiale et
     libre de redevances en vertu des revendications de brevet essentielles de
     ce Contributeur pour fabriquer, utiliser, vendre, offrir à la vente,
     importer et autrement exécuter, modifier et propager le contenu de ce
     Logiciel, dans la mesure permise par les droits accordés aux présentes.

10.2 REPRÉSAILLES EN MATIÈRE DE BREVETS
     Si vous intentez une action en contrefaçon de brevet (incluant une demande
     reconventionnelle ou une mise en cause) contre toute entité alléguant que le
     Logiciel, ou une Contribution incorporée dans le Logiciel, constitue une
     contrefaçon directe ou contributive de brevet, toutes les licences de brevet
     qui vous ont été accordées dans le cadre du présent Contrat pour ce Logiciel
     prennent fin à la date de dépôt de cette action.

10.3 BREVETS DE TECHNOLOGIES BUDGIE
     Rien dans le présent Contrat ne vous accorde de droits sur les brevets,
     demandes de brevets ou secrets commerciaux de Technologies Budgie, autres
     que la licence de brevet limitée décrite à la Section 10.1 concernant les
     Contributions. Les algorithmes propriétaires, les méthodes de compression,
     la logique d'Intégration IA Triage et les techniques d'optimisation DuckDB
     de Technologies Budgie peuvent être soumis à la protection par brevet.
     Aucune licence sur ces brevets n'est accordée dans le cadre du présent
     Contrat, sauf dans la mesure nécessaire pour exercer les droits
     explicitement accordés à la Section 2.

================================================================================
SECTION 11 — TERMINATION / RÉSILIATION
================================================================================

--- ENGLISH ---

11.1 AUTOMATIC TERMINATION FOR BREACH
     Your rights under this Agreement will terminate automatically, without
     notice, if You fail to comply with any term or condition hereof. Upon
     termination:

     (a) All rights granted to You under this Agreement immediately cease;
     (b) You must immediately stop all use, copying, distribution, and
         modification of the Software;
     (c) You must destroy or remove all copies of the Software in Your
         possession or control, including copies in Container registries,
         package repositories, and deployment environments;
     (d) Termination does not relieve You of liability for violations that
         occurred prior to termination.

11.2 CURE PERIOD
     If Your failure to comply is curable, Technologies Budgie may, at its
     discretion, provide written notice identifying the breach and a cure
     period of not less than thirty (30) days. If You cure the breach within
     the specified period, your rights shall be automatically reinstated,
     provided this is the first such cure period granted. Technologies Budgie
     is under no obligation to provide more than one cure period per Licensee.

11.3 TERMINATION FOR SAAS VIOLATION
     Any use of the Software in violation of Section 3.1 (SaaS Use without an
     Enterprise License) shall result in immediate termination of all rights
     under this Agreement, without cure period. Technologies Budgie reserves
     the right to seek immediate injunctive relief and damages for such
     violations.

11.4 SURVIVAL
     The following Sections survive termination of this Agreement: Section 1
     (Definitions), Section 6 (Contributor Terms), Section 7 (Warranties),
     Section 8 (Limitation of Liability), Section 12 (Governing Law), Section
     13 (Dispute Resolution), and Section 14 (Enforcement). Any obligation to
     cease use and destroy copies also survives termination.

11.5 TERMINATION OF DOWNSTREAM LICENSES
     Termination of Your rights under this Agreement does not automatically
     terminate the rights of third parties who have received the Software from
     You under this Agreement, provided they remain in full compliance. However,
     You may not continue to Distribute the Software or support such third
     parties after termination of Your own rights.

--- FRANÇAIS ---

11.1 RÉSILIATION AUTOMATIQUE POUR VIOLATION
     Vos droits dans le cadre du présent Contrat prennent fin automatiquement,
     sans préavis, si vous ne respectez pas l'un ou l'autre des termes ou
     conditions des présentes. Lors de la résiliation :

     (a) Tous les droits qui vous ont été accordés dans le cadre du présent
         Contrat cessent immédiatement ;
     (b) Vous devez immédiatement cesser toute utilisation, copie, distribution
         et modification du Logiciel ;
     (c) Vous devez détruire ou supprimer toutes les copies du Logiciel en votre
         possession ou contrôle, incluant les copies dans les registres de
         conteneurs, les dépôts de paquets et les environnements de déploiement ;
     (d) La résiliation ne vous dégage pas de votre responsabilité pour les
         violations survenues avant la résiliation.

11.2 PÉRIODE DE REMÉDIATION
     Si votre manquement est remédiable, Technologies Budgie peut, à sa
     discrétion, fournir un avis écrit identifiant la violation et une période de
     remédiation d'au moins trente (30) jours. Si vous remédiez à la violation
     dans le délai spécifié, vos droits seront automatiquement rétablis, à
     condition qu'il s'agisse de la première période de remédiation accordée.
     Technologies Budgie n'est pas obligé d'accorder plus d'une période de
     remédiation par Licencié.

11.3 RÉSILIATION POUR VIOLATION SAAS
     Toute utilisation du Logiciel en violation de la Section 3.1 (Usage SaaS
     sans Licence Entreprise) entraîne la résiliation immédiate de tous les
     droits dans le cadre du présent Contrat, sans période de remédiation.
     Technologies Budgie se réserve le droit de demander une injonction
     immédiate et des dommages-intérêts pour de telles violations.

11.4 SURVIE
     Les Sections suivantes survivent à la résiliation du présent Contrat :
     Section 1 (Définitions), Section 6 (Conditions des Contributeurs), Section
     7 (Garanties), Section 8 (Limitation de Responsabilité), Section 12 (Droit
     Applicable), Section 13 (Résolution des Conflits), et Section 14 (Application).
     Toute obligation de cesser l'utilisation et de détruire les copies survit
     également à la résiliation.

11.5 RÉSILIATION DES LICENCES EN AVAL
     La résiliation de vos droits dans le cadre du présent Contrat ne met pas
     automatiquement fin aux droits des tiers qui ont reçu le Logiciel de vous
     dans le cadre du présent Contrat, à condition qu'ils restent pleinement
     conformes. Cependant, vous ne pouvez pas continuer à Distribuer le Logiciel
     ni soutenir ces tiers après la résiliation de vos propres droits.

================================================================================
SECTION 12 — GOVERNING LAW / DROIT APPLICABLE
================================================================================

--- ENGLISH ---

12.1 CHOICE OF LAW
     This Agreement and any dispute, claim, or controversy arising out of or
     relating to it (including its formation, validity, breach, performance, or
     termination) shall be governed by and construed in accordance with the laws
     of the Province of Québec and the applicable federal laws of Canada, without
     regard to conflicts of law principles that would require the application of
     the laws of any other jurisdiction.

12.2 EXCLUSION OF UN CONVENTION
     The United Nations Convention on Contracts for the International Sale of
     Goods (CISG) is expressly excluded and shall not apply to this Agreement.

12.3 CIVIL CODE OF QUÉBEC
     Where applicable, the provisions of the Civil Code of Québec (RLRQ c CCQ-1991)
     shall apply to the interpretation of this Agreement, particularly with
     respect to obligations, contracts, and liability.

12.4 CONSUMER PROTECTION
     If You are a consumer within the meaning of the Consumer Protection Act
     (RLRQ c P-40.1), nothing in this Agreement shall be construed to limit your
     rights as a consumer under that Act. However, the Software is primarily
     designed for professional, technical, and organizational use, and most
     Licensees will not be consumers within the meaning of that Act.

--- FRANÇAIS ---

12.1 CHOIX DU DROIT
     Le présent Contrat et tout litige, réclamation ou controverse en découlant
     ou s'y rapportant (incluant sa formation, sa validité, sa violation, son
     exécution ou sa résiliation) seront régis et interprétés conformément aux
     lois de la Province de Québec et aux lois fédérales applicables du Canada,
     sans égard aux principes de conflit de lois qui exigeraient l'application
     des lois de toute autre juridiction.

12.2 EXCLUSION DE LA CONVENTION DES NATIONS UNIES
     La Convention des Nations Unies sur les contrats de vente internationale de
     marchandises (CVIM) est expressément exclue et ne s'applique pas au présent
     Contrat.

12.3 CODE CIVIL DU QUÉBEC
     Le cas échéant, les dispositions du Code civil du Québec (RLRQ c CCQ-1991)
     s'appliqueront à l'interprétation du présent Contrat, notamment en ce qui
     concerne les obligations, les contrats et la responsabilité.

12.4 PROTECTION DU CONSOMMATEUR
     Si vous êtes un consommateur au sens de la Loi sur la protection du
     consommateur (RLRQ c P-40.1), rien dans le présent Contrat ne sera
     interprété comme limitant vos droits en tant que consommateur en vertu de
     cette loi. Cependant, le Logiciel est principalement conçu pour un usage
     professionnel, technique et organisationnel, et la plupart des Licenciés ne
     seront pas des consommateurs au sens de cette loi.

================================================================================
SECTION 13 — DISPUTE RESOLUTION / RÉSOLUTION DES CONFLITS
================================================================================

--- ENGLISH ---

13.1 MANDATORY NEGOTIATION
     Before initiating any formal legal proceeding, the parties agree to attempt
     to resolve any dispute, claim, or controversy arising out of or relating to
     this Agreement through good-faith negotiation for a period of not less than
     thirty (30) days following written notice from the disputing party to the
     other party identifying the nature of the dispute.

13.2 JURISDICTION
     If negotiation fails to resolve the dispute, the parties irrevocably submit
     to the exclusive jurisdiction of the courts of the Province of Québec,
     judicial district of Montréal (or such other Québec judicial district as
     Technologies Budgie designates at the time of filing), for the resolution
     of any and all disputes arising out of or relating to this Agreement.

13.3 LANGUAGE OF PROCEEDINGS
     Judicial proceedings conducted in Québec may be conducted in French or
     English as determined by the applicable rules of the court and the parties'
     rights under the Charter of the French Language.

13.4 INJUNCTIVE RELIEF
     Notwithstanding the above, Technologies Budgie shall be entitled to seek
     immediate injunctive or other equitable relief in any court of competent
     jurisdiction to prevent irreparable harm arising from a violation or
     threatened violation of this Agreement, without the requirement to post
     bond or other security. The seeking of such relief shall not constitute a
     waiver of any other right or remedy.

13.5 ATTORNEY'S FEES
     In any legal proceeding arising out of this Agreement in which Technologies
     Budgie is the prevailing party, You agree to pay Technologies Budgie's
     reasonable attorney's fees, court costs, and other litigation expenses. This
     provision shall apply in addition to any other remedies available to
     Technologies Budgie.

--- FRANÇAIS ---

13.1 NÉGOCIATION OBLIGATOIRE
     Avant d'engager toute procédure juridique formelle, les parties conviennent
     de tenter de résoudre tout litige, réclamation ou controverse découlant du
     présent Contrat ou s'y rapportant par voie de négociation de bonne foi
     pendant une période d'au moins trente (30) jours suivant un avis écrit de
     la partie en litige à l'autre partie identifiant la nature du litige.

13.2 JURIDICTION
     Si la négociation ne résout pas le litige, les parties se soumettent
     irrévocablement à la compétence exclusive des tribunaux de la Province de
     Québec, district judiciaire de Montréal (ou tel autre district judiciaire du
     Québec que Technologies Budgie désigne au moment du dépôt), pour la
     résolution de tout litige découlant du présent Contrat ou s'y rapportant.

13.3 LANGUE DES PROCÉDURES
     Les procédures judiciaires conduites au Québec peuvent être menées en
     français ou en anglais tel que déterminé par les règles applicables du
     tribunal et les droits des parties en vertu de la Charte de la langue
     française.

13.4 INJONCTION
     Nonobstant ce qui précède, Technologies Budgie aura le droit de demander
     une injonction immédiate ou tout autre redressement en equity devant tout
     tribunal compétent pour prévenir un préjudice irréparable découlant d'une
     violation ou d'une menace de violation du présent Contrat, sans obligation
     de déposer une caution ou autre garantie. La recherche d'un tel redressement
     ne constituera pas une renonciation à tout autre droit ou recours.

13.5 HONORAIRES D'AVOCAT
     Dans toute procédure juridique découlant du présent Contrat dans laquelle
     Technologies Budgie est la partie qui obtient gain de cause, vous acceptez
     de payer les honoraires d'avocat raisonnables de Technologies Budgie, les
     frais de justice et les autres frais de litige. Cette disposition s'applique
     en plus de tout autre recours disponible pour Technologies Budgie.

================================================================================
SECTION 14 — ENFORCEMENT / APPLICATION
================================================================================

--- ENGLISH ---

14.1 MONITORING AND AUDIT
     Technologies Budgie reserves the right, upon reasonable written notice, to
     audit Your use of the Software to verify compliance with this Agreement.
     You agree to cooperate with any such audit and to provide Technologies
     Budgie with reasonable access to relevant records, configurations, and
     deployment information. Audits shall be conducted no more than once per
     calendar year absent evidence of violation.

14.2 VIOLATIONS AND REMEDIES
     Technologies Budgie may pursue any or all of the following remedies for
     violation of this Agreement:

     (a) Termination of Your license under Section 11;
     (b) Injunctive relief to prevent continued or threatened violation;
     (c) Damages, including but not limited to: actual damages, statutory
         damages under applicable copyright law, disgorgement of profits
         attributable to the violation, and reasonable royalties for unlicensed
         SaaS use calculated at no less than the applicable Enterprise Edition
         fee for the period of violation;
     (d) Recovery of attorney's fees and costs as permitted by law;
     (e) Any other remedy available at law or in equity.

14.3 REPORTING VIOLATIONS
     To report suspected violations of this Agreement, contact:
     budgie@mailfence.com

14.4 WAIVER
     Failure by Technologies Budgie to enforce any provision of this Agreement
     at any time shall not be construed as a waiver of Technologies Budgie's
     right to enforce that provision or any other provision in the future.
     No waiver shall be effective unless made in writing and signed by an
     authorized representative of Technologies Budgie.

14.5 SEVERABILITY
     If any provision of this Agreement is held by a court of competent
     jurisdiction to be invalid, illegal, or unenforceable, that provision
     shall be modified to the minimum extent necessary to make it enforceable,
     and the remaining provisions of this Agreement shall continue in full
     force and effect.

14.6 ENTIRE AGREEMENT
     This Agreement constitutes the entire agreement between You and Technologies
     Budgie with respect to the Community Edition of LogDock, and supersedes all
     prior or contemporaneous representations, understandings, agreements, or
     communications, whether written or oral, relating to the subject matter
     hereof. No amendment to this Agreement shall be binding unless made in
     writing and published by Technologies Budgie at
     https://technologiesbudgie.page.dev.

14.7 FUTURE VERSIONS OF THIS LICENSE
     Technologies Budgie may publish revised versions of this License from time
     to time. Each version will be given a distinguishing version number. You may
     use the Software under any previously published version of this License, or
     under the most recently published version. Revised versions will not
     retroactively remove rights granted under prior versions without a
     transition period of no less than twelve (12) months.

--- FRANÇAIS ---

14.1 SURVEILLANCE ET VÉRIFICATION
     Technologies Budgie se réserve le droit, sur avis écrit raisonnable, de
     vérifier votre utilisation du Logiciel pour vérifier la conformité au
     présent Contrat. Vous acceptez de coopérer à toute vérification de ce type
     et de fournir à Technologies Budgie un accès raisonnable aux
     enregistrements, configurations et informations de déploiement pertinents.
     Les vérifications seront conduites au maximum une fois par année civile en
     l'absence de preuve de violation.

14.2 VIOLATIONS ET RECOURS
     Technologies Budgie peut poursuivre tout ou partie des recours suivants pour
     violation du présent Contrat :

     (a) Résiliation de votre licence en vertu de la Section 11 ;
     (b) Injonction pour prévenir la violation continue ou menacée ;
     (c) Dommages-intérêts, incluant notamment : les dommages réels, les
         dommages statutaires en vertu du droit d'auteur applicable, la
         restitution des bénéfices attribuables à la violation, et les redevances
         raisonnables pour l'usage SaaS non licencié calculées à un taux non
         inférieur aux frais applicables de l'Édition Entreprise pour la période
         de violation ;
     (d) Recouvrement des honoraires d'avocat et des frais tels que permis par
         la loi ;
     (e) Tout autre recours disponible en droit ou en equity.

14.3 SIGNALEMENT DES VIOLATIONS
     Pour signaler des violations présumées du présent Contrat, contactez :
     budgie@mailfence.com

14.4 RENONCIATION
     Le défaut de Technologies Budgie d'appliquer toute disposition du présent
     Contrat à tout moment ne sera pas interprété comme une renonciation au droit
     de Technologies Budgie d'appliquer cette disposition ou toute autre
     disposition à l'avenir. Aucune renonciation ne sera effective à moins
     d'être faite par écrit et signée par un représentant autorisé de
     Technologies Budgie.

14.5 DIVISIBILITÉ
     Si une disposition du présent Contrat est jugée invalide, illégale ou
     inapplicable par un tribunal compétent, cette disposition sera modifiée dans
     la mesure minimale nécessaire pour la rendre applicable, et les dispositions
     restantes du présent Contrat continueront à s'appliquer pleinement.

14.6 INTÉGRALITÉ DE L'ACCORD
     Le présent Contrat constitue l'intégralité de l'accord entre Vous et
     Technologies Budgie concernant l'Édition Communautaire de LogDock, et
     remplace toutes les représentations, ententes, accords ou communications
     antérieurs ou contemporains, qu'ils soient écrits ou oraux, relatifs à
     l'objet des présentes. Aucun amendement au présent Contrat ne sera
     contraignant à moins d'être fait par écrit et publié par Technologies
     Budgie à https://technologiesbudgie.page.dev.

14.7 VERSIONS FUTURES DE LA PRÉSENTE LICENCE
     Technologies Budgie peut publier des versions révisées de la présente
     Licence de temps à autre. Chaque version se verra attribuer un numéro de
     version distinctif. Vous pouvez utiliser le Logiciel sous toute version
     précédemment publiée de la présente Licence, ou sous la version la plus
     récemment publiée. Les versions révisées ne retireront pas rétroactivement
     les droits accordés sous les versions précédentes sans une période de
     transition d'au moins douze (12) mois.

================================================================================
SECTION 15 — MISCELLANEOUS / DISPOSITIONS DIVERSES
================================================================================

--- ENGLISH ---

15.1 LANGUAGE
     This Agreement is made in both the English and French languages, both of
     which are equally authoritative. In the event of any inconsistency or
     ambiguity, a Québec court shall apply the French version to the extent
     required by applicable law; in all other jurisdictions, the English version
     shall govern.

15.2 NOTICES
     All notices required or permitted under this Agreement shall be in writing
     and delivered to Technologies Budgie at budgie@mailfence.com. Notices
     delivered by email are effective upon confirmation of receipt. Technologies
     Budgie may provide notices to Licensees through publication on
     https://technologiesbudgie.page.dev or through a notice in the Software's
     release notes.

15.3 FORCE MAJEURE
     Neither party shall be in default by reason of any failure in performance
     of this Agreement if such failure arises out of causes beyond the reasonable
     control and without the fault or negligence of such party, including but
     not limited to natural disasters, war, terrorism, pandemic, government
     action, internet infrastructure failure, and other events beyond reasonable
     control.

15.4 RELATIONSHIP OF PARTIES
     This Agreement does not create a partnership, joint venture, employment,
     franchise, or agency relationship between the parties. Neither party has the
     authority to bind the other party.

15.5 EXPORT COMPLIANCE
     You shall comply with all applicable export laws and regulations in Your
     jurisdiction in connection with Your use and distribution of the Software.

15.6 ACCESSIBILITY
     Technologies Budgie is committed to making LogDock accessible. If you have
     accessibility-related questions or require accommodation, contact us at
     budgie@mailfence.com.

--- FRANÇAIS ---

15.1 LANGUE
     Le présent Contrat est rédigé en anglais et en français, les deux versions
     faisant également foi. En cas d'incohérence ou d'ambiguïté, un tribunal
     québécois appliquera la version française dans la mesure requise par la loi
     applicable ; dans toutes les autres juridictions, la version anglaise
     prévaudra.

15.2 AVIS
     Tous les avis requis ou permis dans le cadre du présent Contrat doivent être
     faits par écrit et transmis à Technologies Budgie à budgie@mailfence.com.
     Les avis transmis par courriel sont effectifs dès confirmation de réception.
     Technologies Budgie peut fournir des avis aux Licenciés par publication sur
     https://technologiesbudgie.page.dev ou par un avis dans les notes de
     publication du Logiciel.

15.3 FORCE MAJEURE
     Aucune des parties ne sera en défaut en raison d'une défaillance
     d'exécution du présent Contrat si cette défaillance découle de causes
     indépendantes du contrôle raisonnable et sans la faute ou la négligence de
     cette partie, incluant notamment les catastrophes naturelles, la guerre, le
     terrorisme, la pandémie, les actions gouvernementales, les défaillances
     d'infrastructure internet et autres événements indépendants du contrôle
     raisonnable.

15.4 RELATION ENTRE LES PARTIES
     Le présent Contrat ne crée pas de partenariat, coentreprise, emploi,
     franchise ou relation d'agence entre les parties. Aucune des parties n'a
     le pouvoir de lier l'autre partie.

15.5 CONFORMITÉ AUX EXPORTATIONS
     Vous devez respecter toutes les lois et réglementations applicables en
     matière d'exportation dans votre juridiction en lien avec votre utilisation
     et distribution du Logiciel.

15.6 ACCESSIBILITÉ
     Technologies Budgie s'engage à rendre LogDock accessible. Si vous avez des
     questions relatives à l'accessibilité ou avez besoin d'une adaptation,
     contactez-nous à budgie@mailfence.com.

================================================================================
END OF LOGDOCK COMMUNITY EDITION LICENSE AGREEMENT
FIN DU CONTRAT DE LICENCE LOGDOCK ÉDITION COMMUNAUTAIRE
================================================================================

Copyright © 2026 Technologies Budgie
Licensed under the LogDock Community Edition License. / Licencé sous LogDock Community Edition License.    
https://technologiesbudgie.page.dev | budgie@mailfence.com