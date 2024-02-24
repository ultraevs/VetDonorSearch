import pickle

from text_processing import main


with open('model.pkl', 'rb') as model_params_file:
    model_params = pickle.load(model_params_file)

with open('vectorizer.pkl', 'rb') as vectorizer_params_file:
    vectorizer_params = pickle.load(vectorizer_params_file)


answer = main(text='Нужен донор в городе москва', model=model_params, vectorizer=vectorizer_params, retrain=False)
print(answer)