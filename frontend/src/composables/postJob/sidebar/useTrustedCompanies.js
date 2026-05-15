import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useTrustedCompanies() {
    const trustedCompaniesData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/sidebar')
            trustedCompaniesData.value = res.data.trustedCompanies
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { trustedCompaniesData, loading, error }
}
