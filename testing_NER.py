import spacy

nlp = spacy.load("judgement_annot")

text = """
Gauhati High Court
Mohesh Munda And 2 Ors vs The State Of Assam on 15 March, 2019
Page No.# 1/18
GAHC010122692015
THE GAUHATI HIGH COURT
(HIGH COURT OF ASSAM, NAGALAND, MIZORAM AND ARUNACHAL PRADESH)
Case No. : CRL.A(J) 98/2015
1:MOHESH MUNDA and 2 ORS.
VERSUS
1:THE STATE OF ASSAM
2:SRI AJIT and RAJIB MUNDA
S/O-BUDHUA MUNDA
VILL-KALAKHOWAGAON
ANANDAPUR T.E.
MURACHUK
P.S.-BOKAKHAT
DIST.-GOLAGHAT
ASSAM
Advocate for the Petitioner : MRS BORGOHAIN
Advocate for the Respondent :
Page No.# 2/18 BEFORE HON'BLE MR. JUSTICE MANASH RANJAN PATHAK HON'BLE MR.
JUSTICE AJIT BORTHAKUR For the appellants : Mr. Sidhartha Borgohain, Ms. B. Devi, Advocates.
For the respondent No. 1 : Mr. Makhan Phukan, Additional Public
Prosecutor, Assam.
For the respondent No. 2 :Did not contest.
Date of Hearing : 10-01-2019.
Date of Judgment : 15-03-2019.
"""

doc = nlp(text)
for ent in doc.ents:
    print(ent.text, ent.start_char, ent.end_char, ent.label_,'\n')