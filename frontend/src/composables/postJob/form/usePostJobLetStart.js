import { ref , onMounted} from 'vue'
import axios from 'axios'

export function usePostJobLetStart(){
    const letStartData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/let_start')
            letStartData.value = res.data
        } catch (e) {
            error.value = e.message 
        
        } finally {
            loading.value = false
        }
    })
    return { letStartData, loading, error}
}