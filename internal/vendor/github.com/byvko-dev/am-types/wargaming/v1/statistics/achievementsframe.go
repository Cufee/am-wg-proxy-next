package statistics

// AchievementsFrame - Represents all achievements of a player
type AchievementsFrame struct {
	ArmorPiercer                int `json:"armorpiercer,omitempty" bson:"armorpiercer,omitempty"`
	MedalFadin                  int `json:"medalfadin,omitempty" bson:"medalfadin,omitempty"`
	MedalCarius                 int `json:"medalcarius,omitempty" bson:"medalcarius,omitempty"`
	MedalEkins                  int `json:"medalekins,omitempty" bson:"medalekins,omitempty"`
	CollectorGuP                int `json:"collectorgup,omitempty" bson:"collectorgup,omitempty"`
	MedalHalonen                int `json:"medalhalonen,omitempty" bson:"medalhalonen,omitempty"`
	HeroesOfRassenay            int `json:"heroesofrassenay,omitempty" bson:"heroesofrassenay,omitempty"`
	FirstVictory                int `json:"firstvictory,omitempty" bson:"firstvictory,omitempty"`
	Defender                    int `json:"defender,omitempty" bson:"defender,omitempty"`
	Creative                    int `json:"creative,omitempty" bson:"creative,omitempty"`
	ESportFinal                 int `json:"esportfinal,omitempty" bson:"esportfinal,omitempty"`
	Supporter                   int `json:"supporter,omitempty" bson:"supporter,omitempty"`
	GoldClanRibbonSEA           int `json:"goldclanribbonsea,omitempty" bson:"goldclanribbonsea,omitempty"`
	PlatinumTwisterMedalSEA     int `json:"platinumtwistermedalsea,omitempty" bson:"platinumtwistermedalsea,omitempty"`
	MedalLehvaslaiho            int `json:"medallehvaslaiho,omitempty" bson:"medallehvaslaiho,omitempty"`
	TankExpert                  int `json:"tankexpert,omitempty" bson:"tankexpert,omitempty"`
	ESportQualification         int `json:"esportqualification,omitempty" bson:"esportqualification,omitempty"`
	MarkI                       int `json:"marki,omitempty" bson:"marki,omitempty"`
	MedalSupremacy              int `json:"medalsupremacy,omitempty" bson:"medalsupremacy,omitempty"`
	ParticipantofWGFest2017     int `json:"participantofwgfest2017,omitempty" bson:"participantofwgfest2017,omitempty"`
	MedalTournamentOffseason1   int `json:"medaltournamentoffseason1,omitempty" bson:"medaltournamentoffseason1,omitempty"`
	JointVictory                int `json:"jointvictory,omitempty" bson:"jointvictory,omitempty"`
	MedalTournamentOffseason2   int `json:"medaltournamentoffseason2,omitempty" bson:"medaltournamentoffseason2,omitempty"`
	MedalTournamentOffseason4   int `json:"medaltournamentoffseason4,omitempty" bson:"medaltournamentoffseason4,omitempty"`
	Sniper                      int `json:"sniper,omitempty" bson:"sniper,omitempty"`
	TitleSniper                 int `json:"titlesniper,omitempty" bson:"titlesniper,omitempty"`
	MedalCrucialContribution    int `json:"medalcrucialcontribution,omitempty" bson:"medalcrucialcontribution,omitempty"`
	Scout                       int `json:"scout,omitempty" bson:"scout,omitempty"`
	GoldTwisterMedalRU          int `json:"goldtwistermedalru,omitempty" bson:"goldtwistermedalru,omitempty"`
	TankExpert3                 int `json:"tankexpert3,omitempty" bson:"tankexpert3,omitempty"`
	TankExpert2                 int `json:"tankexpert2,omitempty" bson:"tankexpert2,omitempty"`
	TankExpert1                 int `json:"tankexpert1,omitempty" bson:"tankexpert1,omitempty"`
	TankExpert0                 int `json:"tankexpert0,omitempty" bson:"tankexpert0,omitempty"`
	MarkOfMastery               int `json:"markofmastery,omitempty" bson:"markofmastery,omitempty"`
	TankExpert6                 int `json:"tankexpert6,omitempty" bson:"tankexpert6,omitempty"`
	TankExpert5                 int `json:"tankexpert5,omitempty" bson:"tankexpert5,omitempty"`
	TankExpert4                 int `json:"tankexpert4,omitempty" bson:"tankexpert4,omitempty"`
	GoldTwisterMedalEU          int `json:"goldtwistermedaleu,omitempty" bson:"goldtwistermedaleu,omitempty"`
	ChristmasTreeLevelUpNY2019  int `json:"christmastreelevelupny2019,omitempty" bson:"christmastreelevelupny2019,omitempty"`
	MedalLavrinenko             int `json:"medallavrinenko,omitempty" bson:"medallavrinenko,omitempty"`
	MedalKolobanov              int `json:"medalkolobanov,omitempty" bson:"medalkolobanov,omitempty"`
	MedalLafayettePool          int `json:"medallafayettepool,omitempty" bson:"medallafayettepool,omitempty"`
	GoldClanRibbonEU            int `json:"goldclanribboneu,omitempty" bson:"goldclanribboneu,omitempty"`
	OlimpicGolden               int `json:"olimpicgolden,omitempty" bson:"olimpicgolden,omitempty"`
	MedalKnispel                int `json:"medalknispel,omitempty" bson:"medalknispel,omitempty"`
	Invader                     int `json:"invader,omitempty" bson:"invader,omitempty"`
	GoldTwisterMedalNA          int `json:"goldtwistermedalna,omitempty" bson:"goldtwistermedalna,omitempty"`
	MechanicEngineer            int `json:"mechanicengineer,omitempty" bson:"mechanicengineer,omitempty"`
	MarkOfMasteryII             int `json:"markofmasteryii,omitempty" bson:"markofmasteryii,omitempty"`
	FirstBlood                  int `json:"firstblood,omitempty" bson:"firstblood,omitempty"`
	MedalKay                    int `json:"medalkay,omitempty" bson:"medalkay,omitempty"`
	MedalOrlik                  int `json:"medalorlik,omitempty" bson:"medalorlik,omitempty"`
	MedalBrothersInArms         int `json:"medalbrothersinarms,omitempty" bson:"medalbrothersinarms,omitempty"`
	MedalAbrams                 int `json:"medalabrams,omitempty" bson:"medalabrams,omitempty"`
	MedalAtgm                   int `json:"medalatgm,omitempty" bson:"medalatgm,omitempty"`
	MainGun                     int `json:"maingun,omitempty" bson:"maingun,omitempty"`
	IronMan                     int `json:"ironman,omitempty" bson:"ironman,omitempty"`
	PlatinumClanRibbonEU        int `json:"platinumclanribboneu,omitempty" bson:"platinumclanribboneu,omitempty"`
	PlatinumClanRibbonSEA       int `json:"platinumclanribbonsea,omitempty" bson:"platinumclanribbonsea,omitempty"`
	Warrior                     int `json:"warrior,omitempty" bson:"warrior,omitempty"`
	GoldClanRibbonRU            int `json:"goldclanribbonru,omitempty" bson:"goldclanribbonru,omitempty"`
	MedalRadleyWalters          int `json:"medalradleywalters,omitempty" bson:"medalradleywalters,omitempty"`
	Raider                      int `json:"raider,omitempty" bson:"raider,omitempty"`
	ParticipantofNewStart       int `json:"participantofnewstart,omitempty" bson:"participantofnewstart,omitempty"`
	DiamondClanRibbon           int `json:"diamondclanribbon,omitempty" bson:"diamondclanribbon,omitempty"`
	MedalBillotte               int `json:"medalbillotte,omitempty" bson:"medalbillotte,omitempty"`
	PlatinumTwisterMedalEU      int `json:"platinumtwistermedaleu,omitempty" bson:"platinumtwistermedaleu,omitempty"`
	Diehard                     int `json:"diehard,omitempty" bson:"diehard,omitempty"`
	MasterofContinents          int `json:"masterofcontinents,omitempty" bson:"masterofcontinents,omitempty"`
	Evileye                     int `json:"evileye,omitempty" bson:"evileye,omitempty"`
	Cadet                       int `json:"cadet,omitempty" bson:"cadet,omitempty"`
	SupremacyHunter             int `json:"supremacyhunter,omitempty" bson:"supremacyhunter,omitempty"`
	ContinentalContender        int `json:"continentalcontender,omitempty" bson:"continentalcontender,omitempty"`
	Steelwall                   int `json:"steelwall,omitempty" bson:"steelwall,omitempty"`
	SupremacyLegend             int `json:"supremacylegend,omitempty" bson:"supremacylegend,omitempty"`
	Punisher                    int `json:"punisher,omitempty" bson:"punisher,omitempty"`
	ESport                      int `json:"esport,omitempty" bson:"esport,omitempty"`
	PlatinumTwisterMark         int `json:"platinumtwistermark,omitempty" bson:"platinumtwistermark,omitempty"`
	GoldClanRibbonNA            int `json:"goldclanribbonna,omitempty" bson:"goldclanribbonna,omitempty"`
	MedalPoppel                 int `json:"medalpoppel,omitempty" bson:"medalpoppel,omitempty"`
	MechanicEngineer6           int `json:"mechanicengineer6,omitempty" bson:"mechanicengineer6,omitempty"`
	MechanicEngineer4           int `json:"mechanicengineer4,omitempty" bson:"mechanicengineer4,omitempty"`
	GoldTwisterMedalSEA         int `json:"goldtwistermedalsea,omitempty" bson:"goldtwistermedalsea,omitempty"`
	MechanicEngineer2           int `json:"mechanicengineer2,omitempty" bson:"mechanicengineer2,omitempty"`
	MechanicEngineer3           int `json:"mechanicengineer3,omitempty" bson:"mechanicengineer3,omitempty"`
	MechanicEngineer0           int `json:"mechanicengineer0,omitempty" bson:"mechanicengineer0,omitempty"`
	MechanicEngineer1           int `json:"mechanicengineer1,omitempty" bson:"mechanicengineer1,omitempty"`
	MechanicEngineer5           int `json:"mechanicengineer5,omitempty" bson:"mechanicengineer5,omitempty"`
	MedalTarczay                int `json:"medaltarczay,omitempty" bson:"medaltarczay,omitempty"`
	Sinai                       int `json:"sinai,omitempty" bson:"sinai,omitempty"`
	PattonValley                int `json:"pattonvalley,omitempty" bson:"pattonvalley,omitempty"`
	MedalDeLanglade             int `json:"medaldelanglade,omitempty" bson:"medaldelanglade,omitempty"`
	DiamondTwisterMedal         int `json:"diamondtwistermedal,omitempty" bson:"diamondtwistermedal,omitempty"`
	Beasthunter                 int `json:"beasthunter,omitempty" bson:"beasthunter,omitempty"`
	SupremacyVeteran            int `json:"supremacyveteran,omitempty" bson:"supremacyveteran,omitempty"`
	Kamikaze                    int `json:"kamikaze,omitempty" bson:"kamikaze,omitempty"`
	OlimpicBronze               int `json:"olimpicbronze,omitempty" bson:"olimpicbronze,omitempty"`
	MedalTournamentOffseason3   int `json:"medaltournamentoffseason3,omitempty" bson:"medaltournamentoffseason3,omitempty"`
	PlatinumClanRibbonRU        int `json:"platinumclanribbonru,omitempty" bson:"platinumclanribbonru,omitempty"`
	MedalOskin                  int `json:"medaloskin,omitempty" bson:"medaloskin,omitempty"`
	Invincible                  int `json:"invincible,omitempty" bson:"invincible,omitempty"`
	PlatinumClanRibbonNA        int `json:"platinumclanribbonna,omitempty" bson:"platinumclanribbonna,omitempty"`
	PlatinumTwisterMedalRU      int `json:"platinumtwistermedalru,omitempty" bson:"platinumtwistermedalru,omitempty"`
	ContinentalViceChampion     int `json:"continentalvicechampion,omitempty" bson:"continentalvicechampion,omitempty"`
	OlimpicSilver               int `json:"olimpicsilver,omitempty" bson:"olimpicsilver,omitempty"`
	MarkOfMasteryI              int `json:"markofmasteryi,omitempty" bson:"markofmasteryi,omitempty"`
	ContinentalCompetitor       int `json:"continentalcompetitor,omitempty" bson:"continentalcompetitor,omitempty"`
	MedalTournamentSummerSeason int `json:"medaltournamentsummerseason,omitempty" bson:"medaltournamentsummerseason,omitempty"`
	Mousebane                   int `json:"mousebane,omitempty" bson:"mousebane,omitempty"`
	MedalBrunoPietro            int `json:"medalbrunopietro,omitempty" bson:"medalbrunopietro,omitempty"`
	MedalTournamentSpringSeason int `json:"medaltournamentspringseason,omitempty" bson:"medaltournamentspringseason,omitempty"`
	GoldTwisterMark             int `json:"goldtwistermark,omitempty" bson:"goldtwistermark,omitempty"`
	CollectorWarhammer          int `json:"collectorwarhammer,omitempty" bson:"collectorwarhammer,omitempty"`
	MarkOfMasteryIII            int `json:"markofmasteryiii,omitempty" bson:"markofmasteryiii,omitempty"`
	MedalLeClerc                int `json:"medalleclerc,omitempty" bson:"medalleclerc,omitempty"`
	MedalTournamentProfessional int `json:"medaltournamentprofessional,omitempty" bson:"medaltournamentprofessional,omitempty"`
	MedalCommunityChampion      int `json:"medalcommunitychampion,omitempty" bson:"medalcommunitychampion,omitempty"`
	DiamondTwisterMark          int `json:"diamondtwistermark,omitempty" bson:"diamondtwistermark,omitempty"`
	PlatinumTwisterMedalNA      int `json:"platinumtwistermedalna,omitempty" bson:"platinumtwistermedalna,omitempty"`
	HandOfDeath                 int `json:"handofdeath,omitempty" bson:"handofdeath,omitempty"`
	Huntsman                    int `json:"huntsman,omitempty" bson:"huntsman,omitempty"`
	Camper                      int `json:"camper,omitempty" bson:"camper,omitempty"`
	MedalNikolas                int `json:"medalnikolas,omitempty" bson:"medalnikolas,omitempty"`
	AndroidTest                 int `json:"androidtest,omitempty" bson:"androidtest,omitempty"`
	Sturdy                      int `json:"sturdy,omitempty" bson:"sturdy,omitempty"`
	MedalTwitch                 int `json:"medaltwitch,omitempty" bson:"medaltwitch,omitempty"`
	MedalWGfestTicket           int `json:"medalwgfestticket,omitempty" bson:"medalwgfestticket,omitempty"`
	ChampionofNewStart          int `json:"championofnewstart,omitempty" bson:"championofnewstart,omitempty"`
}

// Add achievements from b to a
func (a *AchievementsFrame) Add(b *AchievementsFrame) {
	// Achievements
	a.ArmorPiercer = a.ArmorPiercer + b.ArmorPiercer
	a.MedalFadin = a.MedalFadin + b.MedalFadin
	a.MedalCarius = a.MedalCarius + b.MedalCarius
	a.MedalEkins = a.MedalEkins + b.MedalEkins
	a.CollectorGuP = a.CollectorGuP + b.CollectorGuP
	a.MedalHalonen = a.MedalHalonen + b.MedalHalonen
	a.HeroesOfRassenay = a.HeroesOfRassenay + b.HeroesOfRassenay
	a.FirstVictory = a.FirstVictory + b.FirstVictory
	a.Defender = a.Defender + b.Defender
	a.Creative = a.Creative + b.Creative
	a.ESportFinal = a.ESportFinal + b.ESportFinal
	a.Supporter = a.Supporter + b.Supporter
	a.GoldClanRibbonSEA = a.GoldClanRibbonSEA + b.GoldClanRibbonSEA
	a.PlatinumTwisterMedalSEA = a.PlatinumTwisterMedalSEA + b.PlatinumTwisterMedalSEA
	a.MedalLehvaslaiho = a.MedalLehvaslaiho + b.MedalLehvaslaiho
	a.TankExpert = a.TankExpert + b.TankExpert
	a.ESportQualification = a.ESportQualification + b.ESportQualification
	a.MarkI = a.MarkI + b.MarkI
	a.MedalSupremacy = a.MedalSupremacy + b.MedalSupremacy
	a.ParticipantofWGFest2017 = a.ParticipantofWGFest2017 + b.ParticipantofWGFest2017
	a.MedalTournamentOffseason1 = a.MedalTournamentOffseason1 + b.MedalTournamentOffseason1
	a.JointVictory = a.JointVictory + b.JointVictory
	a.MedalTournamentOffseason2 = a.MedalTournamentOffseason2 + b.MedalTournamentOffseason2
	a.MedalTournamentOffseason4 = a.MedalTournamentOffseason4 + b.MedalTournamentOffseason4
	a.Sniper = a.Sniper + b.Sniper
	a.TitleSniper = a.TitleSniper + b.TitleSniper
	a.MedalCrucialContribution = a.MedalCrucialContribution + b.MedalCrucialContribution
	a.Scout = a.Scout + b.Scout
	a.GoldTwisterMedalRU = a.GoldTwisterMedalRU + b.GoldTwisterMedalRU
	a.TankExpert3 = a.TankExpert3 + b.TankExpert3
	a.TankExpert2 = a.TankExpert2 + b.TankExpert2
	a.TankExpert1 = a.TankExpert1 + b.TankExpert1
	a.TankExpert0 = a.TankExpert0 + b.TankExpert0
	a.MarkOfMastery = a.MarkOfMastery + b.MarkOfMastery
	a.TankExpert6 = a.TankExpert6 + b.TankExpert6
	a.TankExpert5 = a.TankExpert5 + b.TankExpert5
	a.TankExpert4 = a.TankExpert4 + b.TankExpert4
	a.GoldTwisterMedalEU = a.GoldTwisterMedalEU + b.GoldTwisterMedalEU
	a.ChristmasTreeLevelUpNY2019 = a.ChristmasTreeLevelUpNY2019 + b.ChristmasTreeLevelUpNY2019
	a.MedalLavrinenko = a.MedalLavrinenko + b.MedalLavrinenko
	a.MedalKolobanov = a.MedalKolobanov + b.MedalKolobanov
	a.MedalLafayettePool = a.MedalLafayettePool + b.MedalLafayettePool
	a.GoldClanRibbonEU = a.GoldClanRibbonEU + b.GoldClanRibbonEU
	a.OlimpicGolden = a.OlimpicGolden + b.OlimpicGolden
	a.MedalKnispel = a.MedalKnispel + b.MedalKnispel
	a.Invader = a.Invader + b.Invader
	a.GoldTwisterMedalNA = a.GoldTwisterMedalNA + b.GoldTwisterMedalNA
	a.MechanicEngineer = a.MechanicEngineer + b.MechanicEngineer
	a.MarkOfMasteryII = a.MarkOfMasteryII + b.MarkOfMasteryII
	a.FirstBlood = a.FirstBlood + b.FirstBlood
	a.MedalKay = a.MedalKay + b.MedalKay
	a.MedalOrlik = a.MedalOrlik + b.MedalOrlik
	a.MedalBrothersInArms = a.MedalBrothersInArms + b.MedalBrothersInArms
	a.MedalAbrams = a.MedalAbrams + b.MedalAbrams
	a.MedalAtgm = a.MedalAtgm + b.MedalAtgm
	a.MainGun = a.MainGun + b.MainGun
	a.IronMan = a.IronMan + b.IronMan
	a.PlatinumClanRibbonEU = a.PlatinumClanRibbonEU + b.PlatinumClanRibbonEU
	a.PlatinumClanRibbonSEA = a.PlatinumClanRibbonSEA + b.PlatinumClanRibbonSEA
	a.Warrior = a.Warrior + b.Warrior
	a.GoldClanRibbonRU = a.GoldClanRibbonRU + b.GoldClanRibbonRU
	a.MedalRadleyWalters = a.MedalRadleyWalters + b.MedalRadleyWalters
	a.Raider = a.Raider + b.Raider
	a.ParticipantofNewStart = a.ParticipantofNewStart + b.ParticipantofNewStart
	a.DiamondClanRibbon = a.DiamondClanRibbon + b.DiamondClanRibbon
	a.MedalBillotte = a.MedalBillotte + b.MedalBillotte
	a.PlatinumTwisterMedalEU = a.PlatinumTwisterMedalEU + b.PlatinumTwisterMedalEU
	a.Diehard = a.Diehard + b.Diehard
	a.MasterofContinents = a.MasterofContinents + b.MasterofContinents
	a.Evileye = a.Evileye + b.Evileye
	a.Cadet = a.Cadet + b.Cadet
	a.SupremacyHunter = a.SupremacyHunter + b.SupremacyHunter
	a.ContinentalContender = a.ContinentalContender + b.ContinentalContender
	a.Steelwall = a.Steelwall + b.Steelwall
	a.SupremacyLegend = a.SupremacyLegend + b.SupremacyLegend
	a.Punisher = a.Punisher + b.Punisher
	a.ESport = a.ESport + b.ESport
	a.PlatinumTwisterMark = a.PlatinumTwisterMark + b.PlatinumTwisterMark
	a.GoldClanRibbonNA = a.GoldClanRibbonNA + b.GoldClanRibbonNA
	a.MedalPoppel = a.MedalPoppel + b.MedalPoppel
	a.MechanicEngineer6 = a.MechanicEngineer6 + b.MechanicEngineer6
	a.MechanicEngineer4 = a.MechanicEngineer4 + b.MechanicEngineer4
	a.GoldTwisterMedalSEA = a.GoldTwisterMedalSEA + b.GoldTwisterMedalSEA
	a.MechanicEngineer2 = a.MechanicEngineer2 + b.MechanicEngineer2
	a.MechanicEngineer3 = a.MechanicEngineer3 + b.MechanicEngineer3
	a.MechanicEngineer0 = a.MechanicEngineer0 + b.MechanicEngineer0
	a.MechanicEngineer1 = a.MechanicEngineer1 + b.MechanicEngineer1
	a.MechanicEngineer5 = a.MechanicEngineer5 + b.MechanicEngineer5
	a.MedalTarczay = a.MedalTarczay + b.MedalTarczay
	a.Sinai = a.Sinai + b.Sinai
	a.PattonValley = a.PattonValley + b.PattonValley
	a.MedalDeLanglade = a.MedalDeLanglade + b.MedalDeLanglade
	a.DiamondTwisterMedal = a.DiamondTwisterMedal + b.DiamondTwisterMedal
	a.Beasthunter = a.Beasthunter + b.Beasthunter
	a.SupremacyVeteran = a.SupremacyVeteran + b.SupremacyVeteran
	a.Kamikaze = a.Kamikaze + b.Kamikaze
	a.OlimpicBronze = a.OlimpicBronze + b.OlimpicBronze
	a.MedalTournamentOffseason3 = a.MedalTournamentOffseason3 + b.MedalTournamentOffseason3
	a.PlatinumClanRibbonRU = a.PlatinumClanRibbonRU + b.PlatinumClanRibbonRU
	a.MedalOskin = a.MedalOskin + b.MedalOskin
	a.Invincible = a.Invincible + b.Invincible
	a.PlatinumClanRibbonNA = a.PlatinumClanRibbonNA + b.PlatinumClanRibbonNA
	a.PlatinumTwisterMedalRU = a.PlatinumTwisterMedalRU + b.PlatinumTwisterMedalRU
	a.ContinentalViceChampion = a.ContinentalViceChampion + b.ContinentalViceChampion
	a.OlimpicSilver = a.OlimpicSilver + b.OlimpicSilver
	a.MarkOfMasteryI = a.MarkOfMasteryI + b.MarkOfMasteryI
	a.ContinentalCompetitor = a.ContinentalCompetitor + b.ContinentalCompetitor
	a.MedalTournamentSummerSeason = a.MedalTournamentSummerSeason + b.MedalTournamentSummerSeason
	a.Mousebane = a.Mousebane + b.Mousebane
	a.MedalBrunoPietro = a.MedalBrunoPietro + b.MedalBrunoPietro
	a.MedalTournamentSpringSeason = a.MedalTournamentSpringSeason + b.MedalTournamentSpringSeason
	a.GoldTwisterMark = a.GoldTwisterMark + b.GoldTwisterMark
	a.CollectorWarhammer = a.CollectorWarhammer + b.CollectorWarhammer
	a.MarkOfMasteryIII = a.MarkOfMasteryIII + b.MarkOfMasteryIII
	a.MedalLeClerc = a.MedalLeClerc + b.MedalLeClerc
	a.MedalTournamentProfessional = a.MedalTournamentProfessional + b.MedalTournamentProfessional
	a.MedalCommunityChampion = a.MedalCommunityChampion + b.MedalCommunityChampion
	a.DiamondTwisterMark = a.DiamondTwisterMark + b.DiamondTwisterMark
	a.PlatinumTwisterMedalNA = a.PlatinumTwisterMedalNA + b.PlatinumTwisterMedalNA
	a.HandOfDeath = a.HandOfDeath + b.HandOfDeath
	a.Huntsman = a.Huntsman + b.Huntsman
	a.Camper = a.Camper + b.Camper
	a.MedalNikolas = a.MedalNikolas + b.MedalNikolas
	a.AndroidTest = a.AndroidTest + b.AndroidTest
	a.Sturdy = a.Sturdy + b.Sturdy
	a.MedalTwitch = a.MedalTwitch + b.MedalTwitch
	a.MedalWGfestTicket = a.MedalWGfestTicket + b.MedalWGfestTicket
	a.ChampionofNewStart = a.ChampionofNewStart + b.ChampionofNewStart
}

// Substracts b from a
func (a *AchievementsFrame) Substract(b *AchievementsFrame) {
	// Achievements
	a.ArmorPiercer = a.ArmorPiercer - b.ArmorPiercer
	a.MedalFadin = a.MedalFadin - b.MedalFadin
	a.MedalCarius = a.MedalCarius - b.MedalCarius
	a.MedalEkins = a.MedalEkins - b.MedalEkins
	a.CollectorGuP = a.CollectorGuP - b.CollectorGuP
	a.MedalHalonen = a.MedalHalonen - b.MedalHalonen
	a.HeroesOfRassenay = a.HeroesOfRassenay - b.HeroesOfRassenay
	a.FirstVictory = a.FirstVictory - b.FirstVictory
	a.Defender = a.Defender - b.Defender
	a.Creative = a.Creative - b.Creative
	a.ESportFinal = a.ESportFinal - b.ESportFinal
	a.Supporter = a.Supporter - b.Supporter
	a.GoldClanRibbonSEA = a.GoldClanRibbonSEA - b.GoldClanRibbonSEA
	a.PlatinumTwisterMedalSEA = a.PlatinumTwisterMedalSEA - b.PlatinumTwisterMedalSEA
	a.MedalLehvaslaiho = a.MedalLehvaslaiho - b.MedalLehvaslaiho
	a.TankExpert = a.TankExpert - b.TankExpert
	a.ESportQualification = a.ESportQualification - b.ESportQualification
	a.MarkI = a.MarkI - b.MarkI
	a.MedalSupremacy = a.MedalSupremacy - b.MedalSupremacy
	a.ParticipantofWGFest2017 = a.ParticipantofWGFest2017 - b.ParticipantofWGFest2017
	a.MedalTournamentOffseason1 = a.MedalTournamentOffseason1 - b.MedalTournamentOffseason1
	a.JointVictory = a.JointVictory - b.JointVictory
	a.MedalTournamentOffseason2 = a.MedalTournamentOffseason2 - b.MedalTournamentOffseason2
	a.MedalTournamentOffseason4 = a.MedalTournamentOffseason4 - b.MedalTournamentOffseason4
	a.Sniper = a.Sniper - b.Sniper
	a.TitleSniper = a.TitleSniper - b.TitleSniper
	a.MedalCrucialContribution = a.MedalCrucialContribution - b.MedalCrucialContribution
	a.Scout = a.Scout - b.Scout
	a.GoldTwisterMedalRU = a.GoldTwisterMedalRU - b.GoldTwisterMedalRU
	a.TankExpert3 = a.TankExpert3 - b.TankExpert3
	a.TankExpert2 = a.TankExpert2 - b.TankExpert2
	a.TankExpert1 = a.TankExpert1 - b.TankExpert1
	a.TankExpert0 = a.TankExpert0 - b.TankExpert0
	a.MarkOfMastery = a.MarkOfMastery - b.MarkOfMastery
	a.TankExpert6 = a.TankExpert6 - b.TankExpert6
	a.TankExpert5 = a.TankExpert5 - b.TankExpert5
	a.TankExpert4 = a.TankExpert4 - b.TankExpert4
	a.GoldTwisterMedalEU = a.GoldTwisterMedalEU - b.GoldTwisterMedalEU
	a.ChristmasTreeLevelUpNY2019 = a.ChristmasTreeLevelUpNY2019 - b.ChristmasTreeLevelUpNY2019
	a.MedalLavrinenko = a.MedalLavrinenko - b.MedalLavrinenko
	a.MedalKolobanov = a.MedalKolobanov - b.MedalKolobanov
	a.MedalLafayettePool = a.MedalLafayettePool - b.MedalLafayettePool
	a.GoldClanRibbonEU = a.GoldClanRibbonEU - b.GoldClanRibbonEU
	a.OlimpicGolden = a.OlimpicGolden - b.OlimpicGolden
	a.MedalKnispel = a.MedalKnispel - b.MedalKnispel
	a.Invader = a.Invader - b.Invader
	a.GoldTwisterMedalNA = a.GoldTwisterMedalNA - b.GoldTwisterMedalNA
	a.MechanicEngineer = a.MechanicEngineer - b.MechanicEngineer
	a.MarkOfMasteryII = a.MarkOfMasteryII - b.MarkOfMasteryII
	a.FirstBlood = a.FirstBlood - b.FirstBlood
	a.MedalKay = a.MedalKay - b.MedalKay
	a.MedalOrlik = a.MedalOrlik - b.MedalOrlik
	a.MedalBrothersInArms = a.MedalBrothersInArms - b.MedalBrothersInArms
	a.MedalAbrams = a.MedalAbrams - b.MedalAbrams
	a.MedalAtgm = a.MedalAtgm - b.MedalAtgm
	a.MainGun = a.MainGun - b.MainGun
	a.IronMan = a.IronMan - b.IronMan
	a.PlatinumClanRibbonEU = a.PlatinumClanRibbonEU - b.PlatinumClanRibbonEU
	a.PlatinumClanRibbonSEA = a.PlatinumClanRibbonSEA - b.PlatinumClanRibbonSEA
	a.Warrior = a.Warrior - b.Warrior
	a.GoldClanRibbonRU = a.GoldClanRibbonRU - b.GoldClanRibbonRU
	a.MedalRadleyWalters = a.MedalRadleyWalters - b.MedalRadleyWalters
	a.Raider = a.Raider - b.Raider
	a.ParticipantofNewStart = a.ParticipantofNewStart - b.ParticipantofNewStart
	a.DiamondClanRibbon = a.DiamondClanRibbon - b.DiamondClanRibbon
	a.MedalBillotte = a.MedalBillotte - b.MedalBillotte
	a.PlatinumTwisterMedalEU = a.PlatinumTwisterMedalEU - b.PlatinumTwisterMedalEU
	a.Diehard = a.Diehard - b.Diehard
	a.MasterofContinents = a.MasterofContinents - b.MasterofContinents
	a.Evileye = a.Evileye - b.Evileye
	a.Cadet = a.Cadet - b.Cadet
	a.SupremacyHunter = a.SupremacyHunter - b.SupremacyHunter
	a.ContinentalContender = a.ContinentalContender - b.ContinentalContender
	a.Steelwall = a.Steelwall - b.Steelwall
	a.SupremacyLegend = a.SupremacyLegend - b.SupremacyLegend
	a.Punisher = a.Punisher - b.Punisher
	a.ESport = a.ESport - b.ESport
	a.PlatinumTwisterMark = a.PlatinumTwisterMark - b.PlatinumTwisterMark
	a.GoldClanRibbonNA = a.GoldClanRibbonNA - b.GoldClanRibbonNA
	a.MedalPoppel = a.MedalPoppel - b.MedalPoppel
	a.MechanicEngineer6 = a.MechanicEngineer6 - b.MechanicEngineer6
	a.MechanicEngineer4 = a.MechanicEngineer4 - b.MechanicEngineer4
	a.GoldTwisterMedalSEA = a.GoldTwisterMedalSEA - b.GoldTwisterMedalSEA
	a.MechanicEngineer2 = a.MechanicEngineer2 - b.MechanicEngineer2
	a.MechanicEngineer3 = a.MechanicEngineer3 - b.MechanicEngineer3
	a.MechanicEngineer0 = a.MechanicEngineer0 - b.MechanicEngineer0
	a.MechanicEngineer1 = a.MechanicEngineer1 - b.MechanicEngineer1
	a.MechanicEngineer5 = a.MechanicEngineer5 - b.MechanicEngineer5
	a.MedalTarczay = a.MedalTarczay - b.MedalTarczay
	a.Sinai = a.Sinai - b.Sinai
	a.PattonValley = a.PattonValley - b.PattonValley
	a.MedalDeLanglade = a.MedalDeLanglade - b.MedalDeLanglade
	a.DiamondTwisterMedal = a.DiamondTwisterMedal - b.DiamondTwisterMedal
	a.Beasthunter = a.Beasthunter - b.Beasthunter
	a.SupremacyVeteran = a.SupremacyVeteran - b.SupremacyVeteran
	a.Kamikaze = a.Kamikaze - b.Kamikaze
	a.OlimpicBronze = a.OlimpicBronze - b.OlimpicBronze
	a.MedalTournamentOffseason3 = a.MedalTournamentOffseason3 - b.MedalTournamentOffseason3
	a.PlatinumClanRibbonRU = a.PlatinumClanRibbonRU - b.PlatinumClanRibbonRU
	a.MedalOskin = a.MedalOskin - b.MedalOskin
	a.Invincible = a.Invincible - b.Invincible
	a.PlatinumClanRibbonNA = a.PlatinumClanRibbonNA - b.PlatinumClanRibbonNA
	a.PlatinumTwisterMedalRU = a.PlatinumTwisterMedalRU - b.PlatinumTwisterMedalRU
	a.ContinentalViceChampion = a.ContinentalViceChampion - b.ContinentalViceChampion
	a.OlimpicSilver = a.OlimpicSilver - b.OlimpicSilver
	a.MarkOfMasteryI = a.MarkOfMasteryI - b.MarkOfMasteryI
	a.ContinentalCompetitor = a.ContinentalCompetitor - b.ContinentalCompetitor
	a.MedalTournamentSummerSeason = a.MedalTournamentSummerSeason - b.MedalTournamentSummerSeason
	a.Mousebane = a.Mousebane - b.Mousebane
	a.MedalBrunoPietro = a.MedalBrunoPietro - b.MedalBrunoPietro
	a.MedalTournamentSpringSeason = a.MedalTournamentSpringSeason - b.MedalTournamentSpringSeason
	a.GoldTwisterMark = a.GoldTwisterMark - b.GoldTwisterMark
	a.CollectorWarhammer = a.CollectorWarhammer - b.CollectorWarhammer
	a.MarkOfMasteryIII = a.MarkOfMasteryIII - b.MarkOfMasteryIII
	a.MedalLeClerc = a.MedalLeClerc - b.MedalLeClerc
	a.MedalTournamentProfessional = a.MedalTournamentProfessional - b.MedalTournamentProfessional
	a.MedalCommunityChampion = a.MedalCommunityChampion - b.MedalCommunityChampion
	a.DiamondTwisterMark = a.DiamondTwisterMark - b.DiamondTwisterMark
	a.PlatinumTwisterMedalNA = a.PlatinumTwisterMedalNA - b.PlatinumTwisterMedalNA
	a.HandOfDeath = a.HandOfDeath - b.HandOfDeath
	a.Huntsman = a.Huntsman - b.Huntsman
	a.Camper = a.Camper - b.Camper
	a.MedalNikolas = a.MedalNikolas - b.MedalNikolas
	a.AndroidTest = a.AndroidTest - b.AndroidTest
	a.Sturdy = a.Sturdy - b.Sturdy
	a.MedalTwitch = a.MedalTwitch - b.MedalTwitch
	a.MedalWGfestTicket = a.MedalWGfestTicket - b.MedalWGfestTicket
	a.ChampionofNewStart = a.ChampionofNewStart - b.ChampionofNewStart
}
