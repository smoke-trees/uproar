import json
from pprint import pprint

with open('lighttag_annotations (4).json') as f:
    data = json.load(f)

#print(len(data['annotations_and_examples']))
print(len(data['annotations_and_examples']))
traindata = []
for i in data['annotations_and_examples']:
    content = i['content']
    annotations = []
    for j in i['annotations']:
        if j['tag'] not in ['PetAdvocate', 'RespAdvocate', 'ThroughLawyer', 'DOJR', 'DOJD', 'Author']:
            annotations.append((j['start'], j['end'],j['tag']))
            print(content[j['start']:j['end']],'...',j['tag'],'\n')

    traindata.append((content, {"entities": annotations}))

print(traindata)
