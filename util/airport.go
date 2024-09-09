package util

import (
	"strings"
)

const airportCodes = "BAZ CAW CCQ FEJ THE POJ VCP AQM NQL ITR PLL PPY CPV CLV RWS JDO BDC BJP PVI FRC CCM PBQ CMG NBV LBR APQ TOW SJK CFB ESI ZFU SLZ BEL TRQ ATM TFL DMT PNB ITB GRU CWB JPR CAU MII SSZ FLB BYO SRA GEL OYK SOD GIG CFC DOU NTM FEC ALT GRP ALQ LAZ TXF UBT UMU GMS GDP NVT RAO PIN BPS UBA PPB SQM RDC AMJ VAG PMW LOI ITI VDC OIA TBT CAC MEA XAP PNG MBK ITE CDJ RSG OAL CPQ JRN JEQ AAI GJM RBB IPU JCM URB APY IGU MNX CKO TGQ PGG APS SQY JCR VIA REC JIA GYN SQX BSS LVB NNU TFF CDI CMP AFL TUR LIP MBZ PHB PLU BVM CKS SFV SSO STZ TEC CMT IOS SSA CLN BCR HRZ BAT BMS FBE NOK IJU OPS RVD CPU UDI PSW SBJ IRE CBW GCV BAU RBR PBB BVS GNM CGH CSS LEC BSB BVB TSQ BPG TUZ VOT LAJ ERN AAX CSW PHI ORX TLZ CSU POO AXE CRQ DIQ SXO AUB PMG BQQ RIA GUZ PBV JTI ERM PIV PAV AQA ZMD SXX CNF POA BVH BGX CQA MCZ MVS FLN IPG JOI CTP APU CTQ CNV BFH AUX PDF JDF MGF STM AJU PIG GVR FEN UNA SJP SWM AAG JPA ITN IPN PTQ AIF OUS MTG RIG SNZ GPB PVH CGR BGV PTO ILB ITP CMC ARS SJL PET GUJ BRA APX OBI FOR MAO ROO BNU MTE LEP NAT LCB PCS ARU MEU MVF NVP JNA PBX TMT CCI ITA JCB VIX MAB RRN IMP CZS LDB CXJ CAF PNZ VAL SEI CIZ CZB CFO MOC SDU BZC ITQ IDO CGB PFB URG CCX BRB MCP DNO"

func IsValidAirport(airportCode string) bool {
	codeList := strings.Split(airportCode, ",")

	for _, code := range codeList {
		found := strings.Contains(airportCodes, code)

		if !found {
			return false
		}
	}

	return true
}
